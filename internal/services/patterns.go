package services

import (
	"errors"
	"fmt"

	"go.mod/internal/analysis"
	"go.mod/internal/domain"
)

func DetectTriangle(pair *domain.Pair) error {
	if len(pair.Candles) < 10 {
		pair.TriangleType = domain.NoTriangle
		return errors.New("not enough candles to detect pattern")
	}

	upperTrend, err := analysis.GetUpperTrendLine(pair.Candles)
	if err != nil {
		pair.TriangleType = domain.NoTriangle
		return fmt.Errorf("failed to find upper trendline")
	}
	pair.UpperTrendLine = upperTrend
	lowerTrend, err := analysis.GetLowerTrendLine(pair.Candles)
	if err != nil {
		pair.TriangleType = domain.NoTriangle
		return fmt.Errorf("failed to find lower trendline")
	}
	pair.LowerTrendLine = lowerTrend
	intersection, err := analysis.IntersectionPoint(upperTrend, lowerTrend)
	if err != nil {
		pair.TriangleType = domain.NoTriangle
		return fmt.Errorf("failed to find itersection point")
	}

	if intersection.X < upperTrend.ClosePoint.X && intersection.X < lowerTrend.ClosePoint.X {
		pair.TriangleType = domain.NoTriangle
		return nil
	}

	switch {
	case lowerTrend.Slope < 0 && upperTrend.Slope < 0:
		pair.TriangleType = domain.Descending
	case lowerTrend.Slope > 0 && upperTrend.Slope > 0:
		pair.TriangleType = domain.Ascending
	default:
		pair.TriangleType = domain.Symmetrical
	}
	return nil
}
