package token

import "errors"

// ErrorClass denotes a category of token error.
type ErrorClass int

const (
	// ErrorClassUnknown denotes an error beyond the scope of this package.
	ErrorClassUnknown ErrorClass = iota

	// ErrorClassInvalidKey denotes an invalid Ed25519 public/private key.
	ErrorClassInvalidKey

	// ErrorClassInvalidTokenData denotes invalid token data.
	ErrorClassInvalidTokenData

	// ErrorClassInvalidTokenSignature denotes an invalid token signature.
	ErrorClassInvalidTokenSignature
)

// Error represents a token related error.
type Error struct {
	class ErrorClass
	error
}

func newError(class ErrorClass, err error) Error {
	return Error{class: class, error: err}
}

// Class reports the ErrorClass of `e`.
func (e Error) Class() ErrorClass { return e.class }

func (e Error) Unwrap() error { return e.error }

// ErrorClassOf is sugar for errors.As(err, *Error), returning the result of Class()
// on success and ErrorClassUnknown otherwise.
func ErrorClassOf(err error) ErrorClass {
	var tokenErr Error
	if errors.As(err, &tokenErr) {
		return tokenErr.Class()
	}
	return ErrorClassUnknown
}
