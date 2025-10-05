package analysis

import (
	"fmt"
	"math"

	"go.mod/internal/domain"
)

//	func (TrendLine TrendLine) getLineEquation() func(x int) float64 {
//		return func(x int) float64 {
//			Y := (TrendLine.Slope * float64(x)) + TrendLine.Y_intercept
//			return Y
//		}
//	}

func GetUpperTrendLine(candles []domain.Candle) (domain.TrendLine, error) {
	if len(candles) < 10 {
		return domain.TrendLine{}, fmt.Errorf("not enough candles")
	}

	maxCloseIndex := findMaxCloseCandleIndex(candles)
	maxClosePoint := domain.Point{
		X: maxCloseIndex,
		Y: candles[maxCloseIndex].Close,
	}

	leftPoint := domain.Point{X: 0, Y: 0}
	rightPoint := domain.Point{X: 0, Y: 0}
	LeftLineSlope := math.Inf(1)
	RightLineSlope := math.Inf(-1)

	for i, candle := range candles {
		if i == maxCloseIndex {
			continue
		}

		s := slope(maxClosePoint.X, maxClosePoint.Y, i, candle.Close)

		if i < maxCloseIndex && s <= LeftLineSlope {
			LeftLineSlope = s
			leftPoint.X = i
			leftPoint.Y = candle.Close
		}
		if i > maxCloseIndex && s >= RightLineSlope {
			RightLineSlope = s
			rightPoint.X = i
			rightPoint.Y = candle.Close
		}

	}

	x1 := maxClosePoint.X - leftPoint.X
	x2 := rightPoint.X - maxClosePoint.X
	// fmt.Println(maxClosePoint)
	// fmt.Println(leftPoint)
	// fmt.Println(rightPoint)
	if x1 > x2 {
		return domain.TrendLine{
			ClosePoint:  maxClosePoint,
			Slope:       LeftLineSlope,
			Y_intercept: maxClosePoint.Y - (LeftLineSlope * float64(maxClosePoint.X)),
		}, nil
	}
	return domain.TrendLine{
		ClosePoint:  maxClosePoint,
		Slope:       RightLineSlope,
		Y_intercept: maxClosePoint.Y - (RightLineSlope * float64(maxClosePoint.X)),
	}, nil

}

func findMaxCloseCandleIndex(candles []domain.Candle) int {
	maxIdx := 0
	maxClose := candles[0].Close

	for i, c := range candles {
		if c.Close > maxClose {
			maxClose = c.Close
			maxIdx = i
		}
	}
	return maxIdx
}
func GetLowerTrendLine(candles []domain.Candle) (domain.TrendLine, error) {
	if len(candles) < 10 {
		return domain.TrendLine{}, fmt.Errorf("not enough candles")
	}

	minCloseIndex := findMinCloseCandleIndex(candles)
	minClosePoint := domain.Point{
		X: minCloseIndex,
		Y: candles[minCloseIndex].Close,
	}

	leftPoint := domain.Point{X: 0, Y: 0}
	rightPoint := domain.Point{X: 0, Y: 0}
	LeftLineSlope := math.Inf(-1)
	RightLineSlope := math.Inf(1)

	for i, candle := range candles {
		if i == minCloseIndex {
			continue
		}

		s := slope(minClosePoint.X, minClosePoint.Y, i, candle.Close)

		if i < minCloseIndex && s >= LeftLineSlope {
			LeftLineSlope = s
			leftPoint.X = i
			leftPoint.Y = candle.Close
		}
		if i > minCloseIndex && s <= RightLineSlope {
			RightLineSlope = s
			rightPoint.X = i
			rightPoint.Y = candle.Close
		}
	}

	x1 := minClosePoint.X - leftPoint.X
	x2 := rightPoint.X - minClosePoint.X

	// fmt.Println(minClosePoint)
	// fmt.Println(leftPoint)
	// fmt.Println(rightPoint)

	if x1 > x2 {
		return domain.TrendLine{
			ClosePoint:  minClosePoint,
			Slope:       LeftLineSlope,
			Y_intercept: minClosePoint.Y - (LeftLineSlope * float64(minClosePoint.X)),
		}, nil
	}
	return domain.TrendLine{
		ClosePoint:  minClosePoint,
		Slope:       RightLineSlope,
		Y_intercept: minClosePoint.Y - (RightLineSlope * float64(minClosePoint.X)),
	}, nil
}

func findMinCloseCandleIndex(candles []domain.Candle) int {
	minIdx := 0
	minClose := candles[0].Close

	for i, c := range candles {
		if c.Close < minClose {
			minClose = c.Close
			minIdx = i
		}
	}

	return minIdx
}

func slope(x1 int, y1 float64, x2 int, y2 float64) float64 {
	if x2 == x1 {
		return math.Inf(1) // prevent division by zero
	}

	return (y2 - y1) / float64(x2-x1)
}

func IntersectionPoint(line1, line2 domain.TrendLine) (domain.Point, error) {
	const epsilon = 1e-6
	if math.Abs(line1.Slope-line2.Slope) < epsilon {
		return domain.Point{}, fmt.Errorf("lines are nearly parallel")
	}

	x := (line2.Y_intercept - line1.Y_intercept) / (line1.Slope - line2.Slope)
	xRounded := roundAwayFromZero(x)
	y := line1.Slope*x + line1.Y_intercept

	return domain.Point{X: xRounded, Y: y}, nil
}
func roundAwayFromZero(x float64) int {
	if x > 0 {
		return int(math.Ceil(x))
	}
	return int(math.Floor(x))
}
