package encoding

import (
	"fmt"

	"github.com/oligarch316/go-token"
	"github.com/oligarch316/go-token/encoding"
	"google.golang.org/grpc/metadata"

	"github.com/oligarch316/go-token/proto/gen/tokenpb"
)

var AuthorizationMeta = HeaderMeta{
	Name:          "authorization",
	ValueEncoding: encoding.PrefixString("Bearer: "),
}

type HeaderMeta struct {
	Name          string
	ValueEncoding token.StringEncoding
}

func (hm HeaderMeta) Encode(t *tokenpb.Token) (metadata.MD, error) {
	s, err := hm.ValueEncoding.Encode(t)
	return metadata.Pairs(hm.Name, s), err
}

func (hm HeaderMeta) Decode(md metadata.MD) (*tokenpb.Token, error) {
	vals := md.Get(hm.Name)

	switch len(vals) {
	case 1:
	case 0:
		return nil, fmt.Errorf("missing '%s' header", hm.Name)
	default:
		return nil, fmt.Errorf("multiple '%s' header values", hm.Name)
	}

	return hm.ValueEncoding.Decode(vals[0])
}
