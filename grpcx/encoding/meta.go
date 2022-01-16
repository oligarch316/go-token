package grpcxenc

import (
	tknx "github.com/oligarch316/go-tokenx"
	tknxenc "github.com/oligarch316/go-tokenx/encoding"
	tknxerr "github.com/oligarch316/go-tokenx/errors"
	"google.golang.org/grpc/metadata"

	"github.com/oligarch316/go-tokenx/proto/gen/tknxpb"
)

const errClass = tknxerr.ClassInvalidTokenData

var AuthorizationMeta = HeaderMeta{
	Name:          "authorization",
	ValueEncoding: tknxenc.PrefixString("Bearer: "),
}

type HeaderMeta struct {
	Name          string
	ValueEncoding tknx.StringEncoding
}

func (hm HeaderMeta) Encode(t *tknxpb.Token) (metadata.MD, error) {
	s, err := hm.ValueEncoding.Encode(t)
	return metadata.Pairs(hm.Name, s), err
}

func (hm HeaderMeta) Decode(md metadata.MD) (*tknxpb.Token, error) {
	vals := md.Get(hm.Name)

	switch len(vals) {
	case 1:
	case 0:
		return nil, tknxerr.Messagef(errClass, "missing '%s' header", hm.Name)
	default:
		return nil, tknxerr.Messagef(errClass, "multiple '%s' header values", hm.Name)
	}

	return hm.ValueEncoding.Decode(vals[0])
}
