package services

import (
	"errors"

	"go.mod/internal/analysis"
	"go.mod/internal/domain"
)

type TriangleType string

const (
	Ascending   TriangleType = "Ascending"
	Descending  TriangleType = "Descending"
	Symmetrical TriangleType = "Symmetrical"
	NoTriangle  TriangleType = "NoTriangle"
)

// DetectTriangle analyzes candle data and returns the type of triangle pattern.
func DetectTriangle(candles []domain.Candle) (TriangleType, error) {
	if len(candles) < 10 {
		return NoTriangle, errors.New("not enough candles to detect pattern")
	}

	upperTrend, err := analysis.GetUpperTrendLine(candles)
	if err != nil {
		return NoTriangle, err
	}

	lowerTrend, err := analysis.GetLowerTrendLine(candles)
	if err != nil {
		return NoTriangle, err
	}

	intersection, err := analysis.IntersectionPoint(upperTrend, lowerTrend)
	if err != nil {
		return NoTriangle, err
	}

	if intersection.X < upperTrend.ClosePoint.X && intersection.X < lowerTrend.ClosePoint.X {
		return NoTriangle, nil
	}

	switch {
	case lowerTrend.Slope < 0 && upperTrend.Slope < 0:
		return Descending, nil
	case lowerTrend.Slope > 0 && upperTrend.Slope > 0:
		return Ascending, nil
	default:
		return Symmetrical, nil
	}
}
