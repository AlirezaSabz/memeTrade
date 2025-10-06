package aggregator

import (
	"fmt"
	"net/http"
	"time"

	"go.mod/internal/domain"
	birdeyeclient "go.mod/internal/infrastructure/birdeye_client"
	moralisclient "go.mod/internal/infrastructure/moralis_client"
)

type Aggregator struct {
	httpClient    *http.Client
	birdeyeClient *birdeyeclient.BirdeyeClient
	moralisClient *moralisclient.MoralisClient
}

func NewAggregator(httpClient *http.Client) *Aggregator {
	m := moralisclient.NewClient(httpClient)
	b := birdeyeclient.NewClient(httpClient)
	return &Aggregator{
		httpClient:    httpClient,
		birdeyeClient: b,
		moralisClient: m,
	}
}

func (c *Aggregator) Tokens() ([]domain.Token, error) {
	tokensMap := make(map[string]*domain.Token)
	trendTokens, err := c.birdeyeClient.TrendingTokens()
	if err != nil {
		return nil, fmt.Errorf("failed to get trending tokens:\n %w", err)
	}
	timeTo := time.Now().Unix()
	timeFrom := timeTo - 2000
	for _, trendToken := range trendTokens {
		pairs, err := c.moralisClient.GetAddressPairs(trendToken.Address)
		if err != nil {
			return nil, fmt.Errorf("failed to get address pairs: %w", err)
		}
		for _, pair := range pairs.Pairs {
			candles, err := c.birdeyeClient.TokenOHLC(pair.PairAddress, timeFrom, timeTo)
			if err != nil {
				return nil, fmt.Errorf("failed to get token ohlc: %w", err)
			}
			_, exists := tokensMap[trendToken.Address]
			if !exists {
				tokensMap[trendToken.Address] = &domain.Token{
					Address: trendToken.Address,
					Pairs:   []domain.Pair{},
				}
			}
			addPairToToken(tokensMap[trendToken.Address], pair, candles, timeFrom, timeTo)
		}
	}

	var tokens []domain.Token
	for _, t := range tokensMap {
		tokens = append(tokens, *t)
	}

	return tokens, nil
}

func (c *Aggregator) FetchNewCandles(tokens *[]domain.Token) error {
	timeTo := time.Now().Unix()
	for _, token := range *tokens {
		for _, pair := range token.Pairs {
			candles, err := c.birdeyeClient.TokenOHLC(pair.Pair, pair.EndTime, timeTo)
			if err != nil {
				return fmt.Errorf("failed to get token ohlc: %w", err)
			}
			pair.Candles = append(pair.Candles, mapCandles(candles)...)
		}
	}
	return nil
}
