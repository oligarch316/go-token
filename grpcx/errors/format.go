package grpcxerr

import (
	tknxerr "github.com/oligarch316/go-tokenx/errors"
	grpcxstatus "github.com/oligarch316/go-tokenx/grpcx/status"
	"google.golang.org/grpc/status"

	"github.com/oligarch316/go-tokenx/proto/gen/grpcxpb"
)

type Formatter func(tknErr tknxerr.Error, msg string) *status.Status

var (
	Short    Formatter = formatShort
	Detailed Formatter = formatDetailed
)

func DetailedFor(classes ...tknxerr.Class) Formatter {
	classMap := make(map[tknxerr.Class]struct{})
	for _, class := range classes {
		classMap[class] = struct{}{}
	}

	return func(tknErr tknxerr.Error, msg string) *status.Status {
		if _, ok := classMap[tknErr.Class]; ok {
			return Detailed(tknErr, msg)
		}
		return Short(tknErr, msg)
	}
}

func formatShort(tknErr tknxerr.Error, msg string) *status.Status {
	return grpcxstatus.New(tknErr.Class, msg)
}

func formatDetailed(tknErr tknxerr.Error, msg string) *status.Status {
	short := grpcxstatus.New(tknErr.Class, msg)
	detailed, err := short.WithDetails(&grpcxpb.ErrorInfo{
		Class: tknErr.Class.String(),
		Cause: tknErr.Error(),
	})

	if err != nil {
		return short
	}

	return detailed
}
