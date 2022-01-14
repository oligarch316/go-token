package token

import (
	"crypto/ed25519"

	"github.com/oligarch316/go-token/errors"
	"google.golang.org/protobuf/proto"

	"github.com/oligarch316/go-token/proto/gen/tokenxpb"
	"google.golang.org/protobuf/types/known/anypb"
)

var (
	errEmptyTokenData        = errors.Message(errors.ClassInvalidTokenData, "empty token data")
	errInvalidTokenSignature = errors.Message(errors.ClassInvalidTokenSignature, "invalid token signature")
	errInvalidPrivateKey     = errors.Message(errors.ClassInvalidKey, "invalid private key")
	errInvalidPublicKey      = errors.Message(errors.ClassInvalidKey, "invalid public key")
)

// StringEncoder defines an encode policy from token to string.
type StringEncoder interface {
	Encode(*tokenxpb.Token) (string, error)
}

// StringDecoder defines a decode policy from string to token.
type StringDecoder interface {
	Decode(string) (*tokenxpb.Token, error)
}

// StringEncoding defines an encode/decode policy between token and string.
type StringEncoding interface {
	StringEncoder
	StringDecoder
}

// Sign marshals `message`, creates an Ed25519 signature for this marshaled data using `privateKey`
// and returns a new token comprised of the results.
// Any error result is compatible with errors.ClassFrom.
func Sign(message proto.Message, privateKey ed25519.PrivateKey) (*tokenxpb.Token, error) {
	if len(privateKey) != ed25519.PrivateKeySize {
		return nil, errInvalidPrivateKey
	}

	data, err := anypb.New(message)
	if err != nil {
		return nil, errors.New(errors.ClassInvalidTokenData, err)
	}

	return &tokenxpb.Token{
		Data:      data,
		Signature: ed25519.Sign(privateKey, data.Value),
	}, nil
}

// Validate verifies the signature of `t` against its data and returns said data.
// Any error result is compatible with errors.ClassFrom.
func Validate(t *tokenxpb.Token, publicKey ed25519.PublicKey) (*anypb.Any, error) {
	if len(publicKey) != ed25519.PublicKeySize {
		return nil, errInvalidPublicKey
	}

	data, signature := t.GetData(), t.GetSignature()
	if data == nil {
		return nil, errEmptyTokenData
	}

	if !ed25519.Verify(publicKey, data.Value, signature) {
		return nil, errInvalidTokenSignature
	}

	return t.Data, nil
}
