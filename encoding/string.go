package encoding

import (
	"encoding/base64"
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/oligarch316/go-token/proto/gen/tokenpb"
)

var URLString = urlString{}

type urlString struct{}

func (urlString) Encode(t *tokenpb.Token) (string, error) {
	b, err := proto.Marshal(t)
	return base64.RawURLEncoding.EncodeToString(b), err
}

func (urlString) Decode(s string) (*tokenpb.Token, error) {
	b, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}

	t := new(tokenpb.Token)
	return t, proto.Unmarshal(b, t)
}

type PrefixString string

func (ps PrefixString) Encode(t *tokenpb.Token) (string, error) {
	s, err := URLString.Encode(t)
	return string(ps) + s, err
}

func (ps PrefixString) Decode(s string) (*tokenpb.Token, error) {
	pLen := len(ps)
	head, tail := s[pLen:], s[:pLen]

	if head != string(ps) {
		return nil, fmt.Errorf("missing '%s' prefix", string(ps))
	}

	return URLString.Decode(tail)
}
