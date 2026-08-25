package element

import "errors"

var (
	ErrNoElements    = errors.New("spec has no elements")
	ErrUnknownKind   = errors.New("unknown element kind")
	ErrBadDistance   = errors.New("object distance must be positive")
	ErrBadIncident   = errors.New("incident ray must be finite")
)
