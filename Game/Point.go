package Game

// ---------------------------------------------------------------------------------------------------------------------
// Constants
// ---------------------------------------------------------------------------------------------------------------------

const (
	xIndex = 0
	yIndex = 1
)

// ---------------------------------------------------------------------------------------------------------------------
// "Point" type definition
// ---------------------------------------------------------------------------------------------------------------------

type Point [2]int

// ---------------------------------------------------------------------------------------------------------------------
// "Point" type methods
// ---------------------------------------------------------------------------------------------------------------------

func (p Point) X() int {
	return p[xIndex]
}

func (p Point) Y() int {
	return p[yIndex]
}

// ---------------------------------------------------------------------------------------------------------------------
// "Point" type constructor
// ---------------------------------------------------------------------------------------------------------------------

func NewPoint(x, y int) Point {
	return Point{x, y}
}
