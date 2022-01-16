package tknxerr

import (
	"errors"
	"fmt"
)

type Error struct {
	Class
	error
}

func (e Error) Unwrap() error { return e.error }

func New(class Class, err error) Error {
	return Error{Class: class, error: err}
}

func Message(class Class, msg string) Error {
	return New(class, errors.New(msg))
}

func Messagef(class Class, format string, a ...interface{}) Error {
	return Message(class, fmt.Sprintf(format, a...))
}

func From(err error) Error {
	var res Error
	if errors.As(err, &res) {
		return res
	}
	return New(ClassUnknown, err)
}

func Classify(err error) Class { return From(err).Class }
