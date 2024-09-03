package domain

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"
)

type SigningSecret struct {
	value string
}

func NewSigningSecret() (SigningSecret, error) {
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return SigningSecret{}, fmt.Errorf("unable to generate random bytes: %v", err)
	}

	return SigningSecret{
		value: fmt.Sprintf("whsec_%s", base64Encoder.WithPadding(base64.NoPadding).EncodeToString(randomBytes)),
	}, nil
}

func (s SigningSecret) String() string {
	return s.value
}

func (s SigningSecret) UnEncoded() ([]byte, error) {
	return base64Encoder.WithPadding(base64.NoPadding).DecodeString(strings.TrimPrefix(s.value, "whsec_"))
}

func (s SigningSecret) IsEmpty() string {
	return s.value
}

func UnMarshallSigningSecret(value string) (SigningSecret, error) {
	//TODO: add proper checks and test this
	return SigningSecret{value: value}, nil
}
