package domain

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	ErrApiKeyIsExpired = errors.New("api key is expired")
	base64Encoder      = base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_.")
)

type ApiKey struct {
	envID     EnvironmentID
	orgID     string
	name      string
	secretKey SecretKey
	createdAt time.Time
	expiresAt *time.Time
}

func NewApiKey(name string, orgID string, envID EnvironmentID, expiresAt *time.Time, isTestApiKey bool) (*ApiKey, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	if orgID == "" {
		return nil, errors.New("orgID cannot be empty")
	}

	if envID == "" {
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
		secretKey: sk,
		name:      name,
		envID:     envID,
		orgID:     orgID,
		createdAt: time.Now(),
		expiresAt: expiresAt,
	}, nil
}

func (a *ApiKey) EnvID() EnvironmentID {
	return a.envID
}

func (a *ApiKey) OrgID() string {
	return a.orgID
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

func (a *ApiKey) IsExpired() bool {
	if a.expiresAt == nil {
		return false
	}

	return a.expiresAt.Before(time.Now())
}

func UnMarshallApiKey(
	envID EnvironmentID,
	orgID,
	name string,
	secretKey SecretKey,
	createdAt time.Time,
	expiresAt *time.Time,
) (*ApiKey, error) {
	return &ApiKey{
		orgID:     orgID,
		envID:     envID,
		name:      name,
		secretKey: secretKey,
		createdAt: createdAt,
		expiresAt: expiresAt,
	}, nil
}

//
// Secret Key
//

type SecretKey struct {
	prefix string
	hash   string
}

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
	return fmt.Sprintf("%s_%s", k.prefix, k.hash)
}

// String returns a trimmed version of the key by showing only the last 5 characters of the hash
func (k SecretKey) String() string {
	return fmt.Sprintf("%s_...%s", k.prefix, k.hash[12:])
}

func UnMarshallSecretKey(value string) (SecretKey, error) {
	chunks := strings.Split(value, "_")

	if len(chunks) < 3 {
		return SecretKey{}, fmt.Errorf("malformed secret key: %s", value)
	}

	var hash string
	if strings.Contains(value, "sbs_live_") {
		hash = strings.Split(value, "sbs_live_")[1]
	}

	if strings.Contains(value, "sbs_test_") {
		hash = strings.Split(value, "sbs_test_")[1]
	}

	return SecretKey{
		prefix: fmt.Sprintf("%s_%s", chunks[0], chunks[1]),
		hash:   hash,
	}, nil
}
