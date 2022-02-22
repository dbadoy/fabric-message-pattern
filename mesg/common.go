package mesg

import (
	"errors"
	"reflect"
)

type CommonRequest struct{}

var (
	ErrInvalidArgment = errors.New("invalid arguments.")
)

func (c CommonRequest) Get() interface{} {
	return c
}

func (c CommonRequest) FieldCount() int {
	v := reflect.ValueOf(c)
	return v.NumField()
}
