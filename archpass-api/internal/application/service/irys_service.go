package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type IrysService struct {
	baseUrl string
	token   string
}

type MetadataInput struct {
	Name        string                   `json:"name"`
	Description string                   `json:"description"`
	Image       string                   `json:"image"`
	Attributes  []map[string]interface{} `json:"attributes"`
	Version     string                   `json:"version"`
	TxId        string                   `json:"txId,omitempty"`
}

func NewIrysService() (*IrysService, error) {
	baseUrl := os.Getenv("IRYS_BASE_URL")
	token := os.Getenv("IRYS_API_TOKEN")

	if baseUrl == "" {
		return nil, fmt.Errorf("IRYS_BASE_URL environment variable is not set")
	}
	if token == "" {
		return nil, fmt.Errorf("IRYS_API_TOKEN environment variable is not set")
	}

	return &IrysService{
		baseUrl: baseUrl,
		token:   token,
	}, nil
}

func (s *IrysService) UploadFile(key string, body *bytes.Buffer, contentType string) (string, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	fw, err := w.CreateFormFile("file", key)
	if err != nil {
		return "", fmt.Errorf("failed to create form file: %v", err)
	}

	if _, err = io.Copy(fw, body); err != nil {
		return "", fmt.Errorf("failed to copy file: %v", err)
	}

	if err := w.WriteField("contentType", contentType); err != nil {
		return "", fmt.Errorf("failed to write content type: %v", err)
	}

	w.Close()
	fmt.Println(fmt.Sprintf("%s/api/upload-image", s.baseUrl))
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/upload-image", s.baseUrl), &b)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("upload failed with status: %d", resp.StatusCode)
	}

	var result struct {
		Success bool   `json:"success"`
		Url     string `json:"url"`
		Error   string `json:"error"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	if !result.Success {
		return "", fmt.Errorf("upload failed: %s", result.Error)
	}

	return result.Url, nil
}

func (s *IrysService) UploadMetadata(metadata MetadataInput) (string, error) {
	jsonData, err := json.Marshal(metadata)
	if err != nil {
		return "", fmt.Errorf("failed to marshal metadata: %v", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/upload-json", s.baseUrl), bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("upload failed with status: %d", resp.StatusCode)
	}

	var result struct {
		Success bool   `json:"success"`
		Url     string `json:"url"`
		Error   string `json:"error"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	if !result.Success {
		return "", fmt.Errorf("upload failed: %s", result.Error)
	}

	return result.Url, nil
}
