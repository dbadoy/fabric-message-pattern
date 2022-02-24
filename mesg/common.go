package mesg

import (
	"errors"
	"reflect"
)

var (
	ErrInvalidArgment = errors.New("invalid arguments.")
)

type CommonRequest struct{}

func (c CommonRequest) Get() interface{} {
	return c
}

func (c CommonRequest) FieldCount() int {
	return reflect.ValueOf(c).NumField()
}
