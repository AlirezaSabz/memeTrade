package domain

type Token struct {
	Address string
	Pairs   []Pair
}
type Pair struct {
	Pair    string
	Candles []Candle
}

type Candle struct {
	UnixTime int64
	High     float64
	Low      float64
	Close    float64
	Open     float64
}
type Point struct {
	X int
	Y float64
}

type TrendLine struct {
	ClosePoint  Point
	Y_intercept float64
	Slope       float64
}
