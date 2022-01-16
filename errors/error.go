package tknxerr

import (
	"errors"
	"fmt"
)

type Error struct {
	class Class
	error
}

func (e Error) Class() Class  { return e.class }
func (e Error) Unwrap() error { return e.error }

func New(class Class, err error) Error {
	return Error{class: class, error: err}
}

func Message(class Class, msg string) Error {
	return New(class, errors.New(msg))
}

func Messagef(class Class, format string, a ...interface{}) Error {
	return Message(class, fmt.Sprintf(format, a...))
}

func ClassFrom(err error) Class {
	var classified Error
	if errors.As(err, &classified) {
		return classified.Class()
	}
	return ClassUnknown
}
