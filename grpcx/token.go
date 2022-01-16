package grpcx

import (
	"context"

	tknxerr "github.com/oligarch316/go-tokenx/errors"
	"google.golang.org/grpc/metadata"

	"github.com/oligarch316/go-tokenx/proto/gen/tknxpb"
)

var errMissingMetadata = tknxerr.Message(tknxerr.ClassInvalidTokenData, "missing metadata")

type MetaEncoder interface {
	Encode(*tknxpb.Token) (metadata.MD, error)
}

type MetaDecoder interface {
	Decode(metadata.MD) (*tknxpb.Token, error)
}

type MetaEncoding interface {
	MetaEncoder
	MetaDecoder
}

func FromIncomingContext(ctx context.Context, dec MetaDecoder) (*tknxpb.Token, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		return dec.Decode(md)
	}

	return nil, errMissingMetadata
}
