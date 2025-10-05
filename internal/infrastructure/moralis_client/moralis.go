package moralisclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type MoralisClient struct {
	httpClient *http.Client
	apiKey     string
}

func NewClient(httpClient *http.Client) *MoralisClient {
	apiKey := os.Getenv("MORALIS_API_KEY")
	if apiKey == "" {
		panic("MORALIS_API_KEY environment variable not set")
	}
	return &MoralisClient{
		httpClient: httpClient,
		apiKey:     apiKey,
	}
}

// GetAddressPairs fetches token pair information for a given token address
func (c *MoralisClient) GetAddressPairs(address string) (*PairResponse, error) {
	url := fmt.Sprintf("https://solana-gateway.moralis.io/token/mainnet/%s/pairs?limit=25", address)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-API-Key", c.apiKey)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status %d: %s", res.StatusCode, string(body))
	}

	var resp PairResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &resp, nil
}
