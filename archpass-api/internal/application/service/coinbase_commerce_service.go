package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"os"
)

type CoinbaseCommerceService struct {
	baseUrl string
	apiKey  string
}

type PriceInput struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type ChargeInput struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	PricingType string            `json:"pricing_type"`
	LocalPrice  PriceInput        `json:"local_price"`
	RedirectURL string            `json:"redirect_url"`
	Metadata    map[string]string `json:"metadata"`
}

type ChargeResponse struct {
	Data struct {
		Code      string    `json:"code"`
		HostedURL string    `json:"hosted_url"`
		ID        uuid.UUID `json:"id"`
		Pricing   struct {
			Local      PriceInput `json:"local"`
			Settlement PriceInput `json:"settlement"`
		} `json:"pricing"`
	} `json:"data"`
	Warnings []string `json:"warnings"`
}

func NewCoinbaseCommerceService() (*CoinbaseCommerceService, error) {
	baseUrl := os.Getenv("COINBASE_COMMERCE_URL")
	apiKey := os.Getenv("COINBASE_COMMERCE_API_KEY")

	if baseUrl == "" {
		return nil, fmt.Errorf("COINBASE_COMMERCE_URL environment variable is not set")
	}
	if apiKey == "" {
		return nil, fmt.Errorf("COINBASE_COMMERCE_API_KEY environment variable is not set")
	}

	return &CoinbaseCommerceService{
		baseUrl: baseUrl,
		apiKey:  apiKey,
	}, nil
}

func (s *CoinbaseCommerceService) CreateCharge(input ChargeInput) (*ChargeResponse, error) {
	jsonData, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal charge input: %v", err)
	}

	fmt.Printf("Creating charge with input: %s\n", string(jsonData))
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/charges", s.baseUrl), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-CC-Api-Key", s.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		var errorResponse struct {
			Error struct {
				Message string `json:"message"`
				Type    string `json:"type"`
			} `json:"error"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {

			return nil, fmt.Errorf("charge creation failed with status: %d, failed to decode error response: %v", resp.StatusCode, err)
		}

		return nil, fmt.Errorf("charge creation failed with status: %d, type: %s, message: %s",
			resp.StatusCode,
			errorResponse.Error.Type,
			errorResponse.Error.Message)
	}

	var result ChargeResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &result, nil
}
