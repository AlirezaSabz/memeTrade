package moralisclient

type PairResponse struct {
	Pairs []Pair `json:"pairs"`
}

type Pair struct {
	ExchangeAddress           string      `json:"exchangeAddress"`
	ExchangeName              string      `json:"exchangeName"`
	ExchangeLogo              string      `json:"exchangeLogo"`
	PairAddress               string      `json:"pairAddress"`
	PairLabel                 string      `json:"pairLabel"`
	UsdPrice                  float64     `json:"usdPrice"`
	UsdPrice24hrPercentChange float64     `json:"usdPrice24hrPercentChange"`
	UsdPrice24hrUsdChange     float64     `json:"usdPrice24hrUsdChange"`
	Volume24hrNative          float64     `json:"volume24hrNative"`
	Volume24hrUsd             float64     `json:"volume24hrUsd"`
	LiquidityUsd              float64     `json:"liquidityUsd"`
	BaseToken                 string      `json:"baseToken"`
	QuoteToken                string      `json:"quoteToken"`
	InactivePair              bool        `json:"inactivePair"`
	PairTokens                []PairToken `json:"pair"`
}

type PairToken struct {
	TokenAddress  string  `json:"tokenAddress"`
	TokenName     string  `json:"tokenName"`
	TokenSymbol   string  `json:"tokenSymbol"`
	TokenLogo     string  `json:"tokenLogo"`
	TokenDecimals string  `json:"tokenDecimals"`
	PairTokenType string  `json:"pairTokenType"`
	LiquidityUsd  float64 `json:"liquidityUsd"`
}
