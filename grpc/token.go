package grpc

import (
	"context"

	"github.com/oligarch316/go-token/errors"
	"github.com/oligarch316/go-token/grpc/status"
	"google.golang.org/grpc/metadata"

	"github.com/oligarch316/go-token/proto/gen/tokenpb"
)

var errMissingMetadata = errors.Message(errors.ClassInvalidTokenData, "missing metadata")

type MetaEncoder interface {
	Encode(*tokenpb.Token) (metadata.MD, error)
}

type MetaDecoder interface {
	Decode(metadata.MD) (*tokenpb.Token, error)
}

type MetaEncoding interface {
	MetaEncoder
	MetaDecoder
}

func FromIncomingContext(ctx context.Context, dec MetaDecoder) (*tokenpb.Token, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		return dec.Decode(md)
	}

	return nil, errMissingMetadata
}

func ConvertError(err error) error { return status.Convert(err).Err() }
