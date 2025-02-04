package dto

type WebhookPayload struct {
	AttemptNumber int    `json:"attempt_number"`
	Event         Event  `json:"event"`
	ID            string `json:"id"`
	ScheduledFor  string `json:"scheduled_for"`
}

type Event struct {
	APIVersion string     `json:"api_version"`
	CreatedAt  string     `json:"created_at"`
	Data       ChargeData `json:"data"`
	ID         string     `json:"id"`
	Resource   string     `json:"resource"`
	Type       string     `json:"type"`
}

type ChargeData struct {
	ID               string     `json:"id"`
	Code             string     `json:"code"`
	Pricing          Pricing    `json:"pricing"`
	Metadata         Metadata   `json:"metadata"`
	Timeline         []Timeline `json:"timeline"`
	Redirects        Redirects  `json:"redirects"`
	Web3Data         Web3Data   `json:"web3_data"`
	CreatedAt        string     `json:"created_at"`
	ExpiresAt        string     `json:"expires_at"`
	HostedURL        string     `json:"hosted_url"`
	BrandColor       string     `json:"brand_color"`
	ChargeKind       string     `json:"charge_kind"`
	ConfirmedAt      string     `json:"confirmed_at"`
	PricingType      string     `json:"pricing_type"`
	SupportEmail     string     `json:"support_email"`
	BrandLogoURL     string     `json:"brand_logo_url"`
	OrganizationName string     `json:"organization_name"`
}

type Pricing struct {
	Local      PricingDetails `json:"local"`
	Settlement PricingDetails `json:"settlement"`
}

type PricingDetails struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Metadata struct {
	OrderId         string `json:"order_id"`
	ContractAddress string `json:"contract_address"`
	TicketId        string `json:"ticket_id"`
}

type Timeline struct {
	Time   string `json:"time"`
	Status string `json:"status"`
}

type Redirects struct {
	CancelURL                string `json:"cancel_url"`
	SuccessURL               string `json:"success_url"`
	WillRedirectAfterSuccess bool   `json:"will_redirect_after_success"`
}

type Web3Data struct {
	FailureEvents     []interface{}      `json:"failure_events"`
	SuccessEvents     []Web3SuccessEvent `json:"success_events"`
	TransferIntent    TransferIntent     `json:"transfer_intent"`
	ContractAddresses map[string]string  `json:"contract_addresses"`
}

type Web3SuccessEvent struct {
	Sender            string `json:"sender"`
	TxHash            string `json:"tx_hsh"`
	Finalized         bool   `json:"finalized"`
	Recipient         string `json:"recipient"`
	Timestamp         string `json:"timestamp"`
	NetworkFeePaid    string `json:"network_fee_paid"`
	InputTokenAmount  string `json:"input_token_amount"`
	InputTokenAddress string `json:"input_token_address"`
}

type TransferIntent struct {
	Metadata TransferMetadata `json:"metadata"`
	CallData CallData         `json:"call_data"`
}

type TransferMetadata struct {
	Sender          string `json:"sender"`
	ChainID         int    `json:"chain_id"`
	ContractAddress string `json:"contract_address"`
}

type CallData struct {
	ID                string `json:"id"`
	Prefix            string `json:"prefix"`
	Deadline          string `json:"deadline"`
	Operator          string `json:"operator"`
	Recipient         string `json:"recipient"`
	Signature         string `json:"signature"`
	FeeAmount         string `json:"fee_amount"`
	RecipientAmount   string `json:"recipient_amount"`
	RecipientCurrency string `json:"recipient_currency"`
	RefundDestination string `json:"refund_destination"`
}

type WebhookResponse struct {
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Data    ResponseData `json:"data"`
}

type ResponseData struct {
	WebhookID   string `json:"webhook_id"`
	ChargeID    string `json:"charge_id"`
	ProcessedAt string `json:"processed_at"`
}
