package birdeyeclient

type BirdeyeResponse struct {
	Data BirdeyeData `json:"data"`
}

type BirdeyeData struct {
	UpdateUnixTime int64   `json:"updateUnixTime"`
	UpdateTime     string  `json:"updateTime"`
	Tokens         []Token `json:"tokens"`
}

type Token struct {
	Address   string  `json:"address"`
	Decimals  int     `json:"decimals"`
	Liquidity float64 `json:"liquidity"`
	LogoURI   string  `json:"logoURI"`
	Name      string  `json:"name"`
	Symbol    string  `json:"symbol"`
	FDV       float64 `json:"fdv"`
	Marketcap float64 `json:"marketcap"`
	Rank      int     `json:"rank"`
	Price     float64 `json:"price"`
	// Volume24hUSD           float64  `json:"volume24hUSD"`
	// Volume24hChangePercent float64  `json:"volume24hChangePercent"`
	// Price24hChangePercent  float64  `json:"price24hChangePercent"`
	// IsScaledUiToken        bool     `json:"isScaledUiToken"`
	// Multiplier             *float64 `json:"multiplier"`
}

type CandlesResponse struct {
	Success bool        `json:"success"`
	Data    CandlesData `json:"data"`
}

type CandlesData struct {
	Items []Candle `json:"items"`
}

type Candle struct {
	Address  string  `json:"address"`
	Close    float64 `json:"c"`
	High     float64 `json:"h"`
	Low      float64 `json:"l"`
	Open     float64 `json:"o"`
	Type     string  `json:"type"`
	UnixTime int64   `json:"unixTime"`
	Volume   float64 `json:"v"`
}
