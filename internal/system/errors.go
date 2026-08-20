package system

import "errors"

var (
	ErrAfocal     = errors.New("afocal system: C=0")
	ErrSingular   = errors.New("no finite solution: denominator zero")
	ErrNoElements = errors.New("system has no elements")
)
