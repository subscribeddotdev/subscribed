package domain

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"
)

type SecretKey struct {
	prefix   string
	hash     string
	checksum uint32
}

var base64Encoder = base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_.")

func newSecretKey(isTestKey bool) (SecretKey, error) {
	prefix := "sbs"
	if isTestKey {
		prefix += "_test"
	} else {
		prefix += "_live"
	}

	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return SecretKey{}, fmt.Errorf("unable to generate random bytes: %v", err)
	}

	return SecretKey{
		prefix: prefix,
		hash:   base64Encoder.WithPadding(base64.NoPadding).EncodeToString(randomBytes),
	}, nil
}

func (k SecretKey) FullKey() string {
	return fmt.Sprintf("%s_%s_%d", k.prefix, k.hash, k.checksum)
}

// String returns a trimmed version of the key by showing only the last 5 characters of the hash
func (k SecretKey) String() string {
	return fmt.Sprintf("%s_...%s_%d", k.prefix, k.hash[12:], k.checksum)
}

type ApiKey struct {
	id        ID
	envID     ID
	name      string
	secretKey SecretKey
	createdAt time.Time
	expiresAt *time.Time
}

func NewApiKey(name string, envID ID, expiresAt *time.Time, isTestApiKey bool) (*ApiKey, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	if envID.IsEmpty() {
		return nil, errors.New("envID cannot be empty")
	}

	if expiresAt != nil && expiresAt.Before(time.Now()) {
		return nil, fmt.Errorf("expiresAt cannot be set in the past")
	}

	sk, err := newSecretKey(isTestApiKey)
	if err != nil {
		return nil, fmt.Errorf("error while creating the secret key: %v", err)
	}

	return &ApiKey{
		id:        NewID(),
		envID:     envID,
		name:      name,
		secretKey: sk,
		createdAt: time.Now().UTC(),
		expiresAt: expiresAt,
	}, nil
}

func (a *ApiKey) Id() ID {
	return a.id
}

func (a *ApiKey) EnvID() ID {
	return a.envID
}

func (a *ApiKey) Name() string {
	return a.name
}

func (a *ApiKey) SecretKey() SecretKey {
	return a.secretKey
}

func (a *ApiKey) CreatedAt() time.Time {
	return a.createdAt
}

func (a *ApiKey) ExpiresAt() *time.Time {
	return a.expiresAt
}
