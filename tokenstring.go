package token

import (
	"crypto/ed25519"

	"github.com/oligarch316/go-token/proto/gen/tokenpb"
	"google.golang.org/protobuf/types/known/anypb"
)

// StringEncoder defines an encode policy from token to string.
type StringEncoder interface {
	Encode(*tokenpb.Token) (string, error)
}

// StringDecoder defines a decode policy from string to token.
type StringDecoder interface {
	Decode(string) (*tokenpb.Token, error)
}

// StringEncoding defines an encode/decode policy between token and string.
type StringEncoding interface {
	StringEncoder
	StringDecoder
}

// ValidateString decodes `s` according to `dec` and calls Validate on the result.
// Any error result is compatible with ErrorClassOf.
func ValidateString(s string, dec StringDecoder, publicKey ed25519.PublicKey) (*anypb.Any, error) {
	t, err := dec.Decode(s)
	if err != nil {
		return nil, newError(ErrorClassInvalidTokenData, err)
	}

	return Validate(t, publicKey)
}
