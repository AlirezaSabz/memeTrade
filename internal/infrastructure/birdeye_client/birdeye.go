package birdeyeclient

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type BirdeyeClient struct {
	httpClient *http.Client
	apiKey     string
}

func NewClient(httpClient *http.Client) *BirdeyeClient {
	apiKey := os.Getenv("BIRDEYE_API_KEY")
	if apiKey == "" {
		log.Fatal("BIRDEYE_API_KEY environment variable not set")
	}
	return &BirdeyeClient{
		httpClient: httpClient,
		apiKey:     apiKey,
	}
}
func (c *BirdeyeClient) TrendingTokens() ([]Token, error) {
	url := "https://public-api.birdeye.so/defi/token_trending?sort_by=rank&sort_type=asc&offset=0&limit=20&ui_amount_mode=scaled"
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-chain", "solana")
	req.Header.Add("X-API-KEY", c.apiKey)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status %d: %s", res.StatusCode, string(body))
	}
	var resp BirdeyeResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	// for _, t := range resp.Data.Tokens {
	// 	fmt.Printf("%s (%s) → price: %.6f, rank: %d\n", t.Name, t.Symbol, t.Price, t.Rank)
	// }
	return resp.Data.Tokens, nil

}

func (c *BirdeyeClient) TokenOHLC(address string, timeFrom int64, timeTo int64) ([]Candle, error) {
	// address = "Czfq3xZZDmsdGdUyrNLtRhGc47cXcZtLG4crryfu44zE"
	url := fmt.Sprintf("https://public-api.birdeye.so/defi/ohlcv/pair?address=%s&type=5m&time_from=%d&time_to=%d", address, timeFrom, timeTo)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-chain", "solana")
	req.Header.Add("X-API-KEY", c.apiKey)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status %d: %s", res.StatusCode, string(body))
	}
	var resp CandlesResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	// fmt.Println(string(body))
	return resp.Data.Items, nil
}
