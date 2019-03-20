package token

import (
	"gopkg.in/square/go-jose.v2/jwt"
)

type Verifier interface {
	// Sign a token and return the serialized cryptographic token.
	Verify(token *string) (*jwt.Claims, error)
}

type joseVerifier struct {
	secret string
}

func NewVerifier(secret string) (Verifier, error) {

	return &joseVerifier{secret: secret}, nil
}

func (j *joseVerifier) Verify(token *string) (*jwt.Claims, error) {
	parsedJWT, err := jwt.ParseSigned(*token)
	if err != nil {
		panic(err)
	}
	resultCl := jwt.Claims{}
	err = parsedJWT.Claims(j.secret, &resultCl)
	return &resultCl, err
}
