package model

import (
	"fmt"
	"time"
)

// Here is game common field limitations, checkers for them and checking error

//----------------------------------------------------------------------------------------------------------------------
// Constants
//----------------------------------------------------------------------------------------------------------------------

const (
	minWidth  = 8
	minHeight = 8
	minMines  = 10
)

// ---------------------------------------------------------------------------------------------------------------------
// "LimitError" type definition
// ---------------------------------------------------------------------------------------------------------------------

type LimitError struct {
	// Name of violated limit
	Violated string
	// Wrong value that was set
	Value int
	// Limit value that must set at least
	Limit int
	// Error time
	When time.Time
}

func (err LimitError) Error() string {
	return fmt.Sprintf(
		"[%s] go_mines/model.LimitError: The %s equals %d is violated limit %d",
		err.When.Format(time.UnixDate),
		err.Violated,
		err.Value,
		err.Limit,
	)
}

func NewLimitError(violated string, value, limit int) LimitError {
	return LimitError{
		Violated: violated,
		Value:    value,
		Limit:    limit,
		When:     time.Now(),
	}
}

// ---------------------------------------------------------------------------------------------------------------------
// Checkers
// ---------------------------------------------------------------------------------------------------------------------

func checkWidth(width int) error {
	if width < minWidth {
		return NewLimitError("width", width, minWidth)
	}

	return nil
}

func checkHeight(height int) error {
	if height < minHeight {
		return NewLimitError("height", height, minHeight)
	}

	return nil
}

func checkMines(mines int) error {
	if mines < minMines {
		return NewLimitError("mines", mines, minMines)
	}

	return nil
}
