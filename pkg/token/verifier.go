package token

import (
	"gopkg.in/square/go-jose.v2/jwt"
)

type Verifier interface {
	// Sign a token and return the serialized cryptographic token.
	Verify(token *string) (*string, error)
}

type joseVerifier struct {
}

func NewJoseVerifier(secret string) (Verifier, error) {

	return &joseVerifier{}, nil
}

func (j *joseVerifier) Verify(token *string) (*string, error) {
	jsonwt, err := jwt.ParseSigned(*token)
	if err != nil {
		panic(err)
	}
	return nil, nil
}
