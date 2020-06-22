package model

import (
	"fmt"
	"time"
)

// -----------------------------------------------------------------------------
// Constants
// -----------------------------------------------------------------------------

const (
	FIELD_CREATION_ERROR_CODE = 1
	FIELD_CREATION_ERROR_NAME = "FieldCreationError"
	CELL_OPEN_ERROR_CODE      = 2
	CELL_OPEN_ERROR_NAME      = "CellOpenError"
)

// -----------------------------------------------------------------------------
// "AppModelError" type definition
// -----------------------------------------------------------------------------

type AppModelError struct {
	Code    int
	Name    string
	Message string
	When    time.Time
}

func (err AppModelError) Error() string {
	return fmt.Sprintf(
		"[%s] App Model Error #%d \"%s\": %s",
		err.When.Format(time.UnixDate),
		err.Code,
		err.Name,
		err.Message,
	)
}

func NewAppModelError(code int, name, message string) AppModelError {
	return AppModelError{code, name, message, time.Now()}
}

func NewFieldCreationError(reason string) AppModelError {
	return NewAppModelError(FIELD_CREATION_ERROR_CODE, FIELD_CREATION_ERROR_NAME, reason)
}

func NewCellOpenError(reason string) AppModelError {
	return NewAppModelError(CELL_OPEN_ERROR_CODE, CELL_OPEN_ERROR_NAME, reason)
}
