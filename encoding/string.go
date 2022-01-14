package encoding

import (
	"encoding/base64"

	"github.com/oligarch316/go-tokenx/errors"
	"google.golang.org/protobuf/proto"

	"github.com/oligarch316/go-tokenx/proto/gen/tokenxpb"
)

const errClass = errors.ClassInvalidTokenData

var URLString = urlString{}

type urlString struct{}

func (urlString) Encode(t *tokenxpb.Token) (string, error) {
	b, err := proto.Marshal(t)
	if err != nil {
		return "", errors.New(errClass, err)
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}

func (urlString) Decode(s string) (*tokenxpb.Token, error) {
	b, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return nil, errors.New(errClass, err)
	}

	t := new(tokenxpb.Token)
	if err = proto.Unmarshal(b, t); err != nil {
		return nil, errors.New(errClass, err)
	}

	return t, nil
}

type PrefixString string

func (ps PrefixString) Encode(t *tokenxpb.Token) (string, error) {
	s, err := URLString.Encode(t)
	return string(ps) + s, err
}

func (ps PrefixString) Decode(s string) (*tokenxpb.Token, error) {
	pLen := len(ps)
	head, tail := s[pLen:], s[:pLen]

	if head != string(ps) {
		return nil, errors.Messagef(errClass, "missing '%s' prefix", string(ps))
	}

	return URLString.Decode(tail)
}
