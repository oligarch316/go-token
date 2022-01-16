package grpcxstatus

import (
	tknxerr "github.com/oligarch316/go-tokenx/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const unknownClassCode = codes.Unknown

var knownClassCodes = map[tknxerr.Class]codes.Code{
	tknxerr.ClassUnknown:               codes.Internal,
	tknxerr.ClassInvalidKey:            codes.Internal,
	tknxerr.ClassInvalidTokenData:      codes.Unauthenticated,
	tknxerr.ClassInvalidTokenSignature: codes.Unauthenticated,
}

func Code(class tknxerr.Class) codes.Code {
	if code, ok := knownClassCodes[class]; ok {
		return code
	}
	return unknownClassCode
}

func New(class tknxerr.Class, msg string) *status.Status {
	return status.New(Code(class), msg)
}

func Newf(class tknxerr.Class, format string, a ...interface{}) *status.Status {
	return status.Newf(Code(class), format, a...)
}
