package aggregator

import (
	"go.mod/internal/domain"
	birdeyeclient "go.mod/internal/infrastructure/birdeye_client"
	moralisclient "go.mod/internal/infrastructure/moralis_client"
)

func addPairToToken(token *domain.Token, pair moralisclient.Pair, candles []birdeyeclient.Candle, startTime int64, endTime int64) {
	token.Pairs = append(token.Pairs, domain.Pair{
		Pair:      pair.PairAddress,
		Candles:   mapCandles(candles),
		StartTime: startTime,
		EndTime:   endTime,
	})
}

func mapCandles(candles []birdeyeclient.Candle) []domain.Candle {
	domainCandles := make([]domain.Candle, 0, len(candles))
	for _, c := range candles {
		domainCandles = append(domainCandles, domain.Candle{
			High:     c.High,
			Low:      c.Low,
			Open:     c.Open,
			Close:    c.Close,
			UnixTime: c.UnixTime,
		})
	}
	return domainCandles
}
