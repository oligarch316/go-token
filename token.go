package token

import (
	"crypto/ed25519"
	"errors"

	"google.golang.org/protobuf/proto"

	"github.com/oligarch316/go-token/proto/gen/tokenpb"
	"google.golang.org/protobuf/types/known/anypb"
)

var (
	errEmptyTokenData        = errors.New("empty token data")
	errInvalidTokenSignature = errors.New("invalid token signature")
	errInvalidPrivateKey     = errors.New("invalid private key")
	errInvalidPublicKey      = errors.New("invalid public key")
)

// Sign marshals `message`, creates an Ed25519 signature for this marshaled data using `privateKey`
// and returns a new token comprised of the results.
// Any error result is compatible with ErrorClassOf.
func Sign(message proto.Message, privateKey ed25519.PrivateKey) (*tokenpb.Token, error) {
	if len(privateKey) != ed25519.PrivateKeySize {
		return nil, newError(ErrorClassInvalidKey, errInvalidPrivateKey)
	}

	data, err := anypb.New(message)
	if err != nil {
		return nil, newError(ErrorClassInvalidTokenData, err)
	}

	return &tokenpb.Token{
		Data:      data,
		Signature: ed25519.Sign(privateKey, data.Value),
	}, nil
}

// Validate verifies the signature of `t` against its data and returns said data.
// Any error result is compatible with ErrorClassOf.
func Validate(t *tokenpb.Token, publicKey ed25519.PublicKey) (*anypb.Any, error) {
	if len(publicKey) != ed25519.PublicKeySize {
		return nil, newError(ErrorClassInvalidKey, errInvalidPublicKey)
	}

	data, signature := t.GetData(), t.GetSignature()
	if data == nil {
		return nil, newError(ErrorClassInvalidTokenData, errEmptyTokenData)
	}

	if !ed25519.Verify(publicKey, data.Value, signature) {
		return nil, newError(ErrorClassInvalidTokenSignature, errInvalidTokenSignature)
	}

	return t.Data, nil
}
