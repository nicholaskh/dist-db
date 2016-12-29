package cmd

import (
	"errors"
)

var (
	ErrSyntaxError   = errors.New("ERR syntax error")
	ErrKeyNotChanged = errors.New("ERR key not changed")
	ErrKeyNotFound   = errors.New("ERR key not found")
)
