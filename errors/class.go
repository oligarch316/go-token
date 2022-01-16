package tknxerr

type Class uint8

const (
	ClassUnknown Class = iota
	ClassInvalidKey
	ClassInvalidTokenData
	ClassInvalidTokenSignature
)

const unknownClassString = "unknown error"

var knownClassStrings = map[Class]string{
	ClassInvalidKey:            "invalid key",
	ClassInvalidTokenData:      "invalid token data",
	ClassInvalidTokenSignature: "invalid token signature",
}

func (c Class) String() string {
	if str, ok := knownClassStrings[c]; ok {
		return str
	}
	return unknownClassString
}
