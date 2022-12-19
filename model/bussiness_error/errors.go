package bussiness_error

import (
	"errors"
)

var ErrNotFound = errors.New("record not found")

var ErrInvalidLineFormat = errors.New("line has invalid line")
