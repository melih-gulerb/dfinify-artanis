package clients

import (
	"artanis/src/models/clients"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type DivineShieldClient struct {
	BaseURL    string
	HttpClient *http.Client
}

type AuthResponse struct {
	Token string       `json:"token"`
	User  clients.User `json:"user"`
}

type AuthRequest struct {
	Token string `json:"token"`
}

func NewDivineShieldClient(baseURL string) *DivineShieldClient {
	return &DivineShieldClient{
		BaseURL: baseURL,
		HttpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *DivineShieldClient) Authorize(token string) (*clients.User, error) {
	authReq := AuthRequest{
		Token: token,
	}

	reqBody, err := json.Marshal(authReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal auth request: %w", err)
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/authoriation", c.BaseURL),
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("divine-shield service returned non-OK status: %d", resp.StatusCode)
	}

	var authResp AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &authResp.User, nil
}
