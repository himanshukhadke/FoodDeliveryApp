package entities

import "math"

type Location [2]float64

func NewLocation(coordinates [2]float64) Location {
	return Location{coordinates[0], coordinates[1]}
}

func (l Location) GetCoordinates() [2]float64 {
	return [2]float64{l[0], l[1]}
}

func (l Location) CalculateDistance(l2 Location) float64 {
	x1, y1 := l[0], l[1]
    x2, y2 := l2[0], l2[1]
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
