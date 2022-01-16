package tknxenc

import (
	"encoding/base64"

	tknxerr "github.com/oligarch316/go-tokenx/errors"
	"google.golang.org/protobuf/proto"

	"github.com/oligarch316/go-tokenx/proto/gen/tknxpb"
)

const errClass = tknxerr.ClassInvalidTokenData

var URLString = urlString{}

type urlString struct{}

func (urlString) Encode(t *tknxpb.Token) (string, error) {
	b, err := proto.Marshal(t)
	if err != nil {
		return "", tknxerr.New(errClass, err)
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}

func (urlString) Decode(s string) (*tknxpb.Token, error) {
	b, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return nil, tknxerr.New(errClass, err)
	}

	t := new(tknxpb.Token)
	if err = proto.Unmarshal(b, t); err != nil {
		return nil, tknxerr.New(errClass, err)
	}

	return t, nil
}

type PrefixString string

func (ps PrefixString) Encode(t *tknxpb.Token) (string, error) {
	s, err := URLString.Encode(t)
	return string(ps) + s, err
}

func (ps PrefixString) Decode(s string) (*tknxpb.Token, error) {
	pLen := len(ps)
	head, tail := s[pLen:], s[:pLen]

	if head != string(ps) {
		return nil, tknxerr.Messagef(errClass, "missing '%s' prefix", string(ps))
	}

	return URLString.Decode(tail)
}
