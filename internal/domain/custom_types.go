package domain

type TriangleType string

const (
	Ascending   TriangleType = "Ascending"
	Descending  TriangleType = "Descending"
	Symmetrical TriangleType = "Symmetrical"
	NoTriangle  TriangleType = "NoTriangle"
)
