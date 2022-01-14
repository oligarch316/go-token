package grpcx

import (
	"context"

	"github.com/oligarch316/go-tokenx/errors"
	"github.com/oligarch316/go-tokenx/grpcx/status"
	"google.golang.org/grpc/metadata"

	"github.com/oligarch316/go-tokenx/proto/gen/tokenxpb"
)

var errMissingMetadata = errors.Message(errors.ClassInvalidTokenData, "missing metadata")

type MetaEncoder interface {
	Encode(*tokenxpb.Token) (metadata.MD, error)
}

type MetaDecoder interface {
	Decode(metadata.MD) (*tokenxpb.Token, error)
}

type MetaEncoding interface {
	MetaEncoder
	MetaDecoder
}

func FromIncomingContext(ctx context.Context, dec MetaDecoder) (*tokenxpb.Token, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		return dec.Decode(md)
	}

	return nil, errMissingMetadata
}

func ConvertError(err error) error { return status.Convert(err).Err() }
