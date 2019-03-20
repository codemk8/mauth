package token

import (
	"fmt"

	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

type Signer interface {
	// Sign a token and return the serialized cryptographic token.
	Sign(token *string) (*string, error)
}

type joseSigner struct {
	signer jose.Signer
}

func NewSigner(secret string) (Signer, error) {
	// Instantiate an encrypter using RSA-OAEP with AES128-GCM. An error would
	// indicate that the selected algorithm(s) are not currently supported.
	//publicKey := &privateKey.PublicKey
	key := []byte("secret")
	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: key}, (&jose.SignerOptions{}).WithType("JWT"))

	if err != nil {
		panic(err)
	}

	cl := jwt.Claims{
		Subject: "subject",
		Issuer:  "issuer",
	}
	raw, err := jwt.Signed(sig).Claims(cl).CompactSerialize()
	if err != nil {
		panic(err)
	}

	fmt.Println(raw)

	return &joseSigner{signer: sig}, nil
}

func (j *joseSigner) Sign(token *string) (*string, error) {
	cl := jwt.Claims{
		Subject: "subject",
		Issuer:  "issuer",
	}
	raw, err := jwt.Signed(j.signer).Claims(cl).CompactSerialize()
	if err != nil {
		panic(err)
	}
	return &raw, nil
}
