package grpc

import (
	"crypto/ed25519"

	"github.com/oligarch316/go-token"
	"github.com/oligarch316/go-token/proto/gen/tokenpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func Sign(message proto.Message, privateKey ed25519.PrivateKey) (*tokenpb.Token, error) {
	res, err := token.Sign(message, privateKey)
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
	}

	return res, err
}

func Validate(t *tokenpb.Token, publicKey ed25519.PublicKey) (*anypb.Any, error) {
	res, err := token.Validate(t, publicKey)
	if err != nil {
		switch token.ErrorClassOf(err) {
		case token.ErrorClassInvalidTokenData, token.ErrorClassInvalidTokenSignature:
			err = status.Error(codes.Unauthenticated, err.Error())
		default:
			err = status.Error(codes.Internal, err.Error())
		}
	}

	return res, err
}
