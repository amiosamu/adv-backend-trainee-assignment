package repoerrors

import "errors"

var (
	ErrNotFound  = errors.New("not found")
	CannotCreate = errors.New("cannot create")
)
