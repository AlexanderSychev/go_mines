package model

import (
	"fmt"
	"math/rand"
	"time"
)

//------------------------------------------------------------------------------
// "Coordinates" type definition
//------------------------------------------------------------------------------

type Coordinates [2]int

// Methods

func (c Coordinates) X() int {
	return c[0]
}

func (c Coordinates) Y() int {
	return c[1]
}

func (c Coordinates) String() string {
	return fmt.Sprintf("Coordinates(x=%d, y=%d)", c.X(), c.Y())
}

// Constructor

func NewCoordinates(x, y int) Coordinates {
	return Coordinates{x, y}
}

//------------------------------------------------------------------------------
// "UniqueCoordinatesSet" type definition
//------------------------------------------------------------------------------

// List of unique coordinates
type UniqueCoordinatesSet struct {
	items []Coordinates
}

// Methods

// Returns "true" if set already contains coordinates with same "x" and "y"
func (set *UniqueCoordinatesSet) Contains(c Coordinates) bool {
	for _, item := range set.items {
		if item.X() == c.X() && item.Y() == c.Y() {
			return true
		}
	}
	return false
}

func (set UniqueCoordinatesSet) Length() int {
	return len(set.items)
}

func (set *UniqueCoordinatesSet) Get(index int) Coordinates {
	return set.items[index]
}

func (set *UniqueCoordinatesSet) Add(c Coordinates) bool {
	contains := set.Contains(c)
	if !contains {
		set.items = append(set.items, c)
	}
	return !contains
}

// Constructors

func NewUniqueCoordinatesSet() UniqueCoordinatesSet {
	return UniqueCoordinatesSet{make([]Coordinates, 0)}
}

func RandomUniqueCoordinatesSet(width, height, mines int) UniqueCoordinatesSet {
	rand.Seed(time.Now().UTC().UnixNano())

	set := NewUniqueCoordinatesSet()

	for i := 0; i < mines; i++ {
		for {
			x := rand.Intn(height)
			y := rand.Intn(width)
			if set.Add(NewCoordinates(x, y)) {
				break
			}
		}
	}

	return set
}
