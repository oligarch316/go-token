package grpcxerr

import (
	"fmt"

	tknxerr "github.com/oligarch316/go-tokenx/errors"
	"google.golang.org/grpc/status"
)

type Error struct {
	tknErr tknxerr.Error
	msg    string
	format Formatter
}

func (e Error) Error() string              { return e.tknErr.Error() }
func (e Error) Unwrap() error              { return e.tknErr }
func (e Error) GRPCStatus() *status.Status { return e.format(e.tknErr, e.msg) }

func Wrap(f Formatter, err error, msg string) Error {
	return Error{tknErr: tknxerr.From(err), msg: msg, format: f}
}

func Wrapf(f Formatter, err error, format string, a ...interface{}) Error {
	return Wrap(f, err, fmt.Sprintf(format, a...))
}

func Convert(f Formatter, err error) Error {
	tknErr := tknxerr.From(err)
	return Error{tknErr: tknErr, msg: tknErr.Class.String(), format: f}
}
