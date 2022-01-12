package grpc

import (
	"context"
	"crypto/ed25519"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/oligarch316/go-token/proto/gen/tokenpb"
	"google.golang.org/protobuf/types/known/anypb"
)

const errMsgMissingMetadata = "missing metadata"

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

func ValidateRequestMeta(md metadata.MD, dec MetaDecoder, publicKey ed25519.PublicKey) (*anypb.Any, error) {
	t, err := dec.Decode(md)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	return Validate(t, publicKey)
}

func ValidateRequestContext(ctx context.Context, dec MetaDecoder, publicKey ed25519.PublicKey) (*anypb.Any, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, errMsgMissingMetadata)
	}

	return ValidateRequestMeta(md, dec, publicKey)
}
