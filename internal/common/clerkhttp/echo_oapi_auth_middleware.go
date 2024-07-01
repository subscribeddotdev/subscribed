package clerkhttp

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/jwks"
	"github.com/clerk/clerk-sdk-go/v2/jwt"
	"github.com/friendsofgo/errors"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	oapimiddleware "github.com/oapi-codegen/echo-middleware"
)

const ctxJwtUserClaims = "clerk_jwt_user_claims"

// EchoOapiAuthMiddleware Implemented based on https://github.com/clerk/clerk-sdk-go/blob/v2/http/middleware.go
// with the goal of working with OAPI Codegen authenticator helper:
// https://github.com/oapi-codegen/oapi-codegen/blob/main/examples/authenticated-api/echo/server/server.go#L37
type EchoOapiAuthMiddleware struct {
}

func NewEchoOapiAuthMiddleware(clerkSecretKey string) *EchoOapiAuthMiddleware {
	// clerk.SetKey(clerkSecretKey)
	return &EchoOapiAuthMiddleware{}
}

func (e *EchoOapiAuthMiddleware) Middleware() openapi3filter.AuthenticationFunc {
	return func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
		r := input.RequestValidationInput.Request
		params := &AuthorizationParams{}

		params.Clock = clerk.NewClock()

		authorization := strings.TrimSpace(r.Header.Get("Authorization"))
		if authorization == "" {
			return errors.New("authorization header cannot be empty")
		}

		token := strings.TrimPrefix(authorization, "Bearer ")
		decoded, err := jwt.Decode(r.Context(), &jwt.DecodeParams{Token: token})
		if err != nil {
			return fmt.Errorf("unable to decode the token: %v", err)
		}

		params.JWK, err = getJWK(r.Context(), params.JWKSClient, decoded.KeyID, params.Clock)
		if err != nil {
			return err
		}

		params.Token = token
		claims, err := jwt.Verify(r.Context(), &params.VerifyParams)
		if err != nil {
			return err
		}

		// Token was verified. Add the session claims to the request context.
		echoCtx := oapimiddleware.GetEchoContext(ctx)
		echoCtx.Set(ctxJwtUserClaims, claims)

		return nil
	}
}

// Retrieve the JSON web key for the provided token from the JWKS set.
// Tries a cached value first, but if there's no value or the entry
// has expired, it will fetch the JWK set from the API and cache the
// value.
func getJWK(ctx context.Context, jwksClient *jwks.Client, kid string, clock clerk.Clock) (*clerk.JSONWebKey, error) {
	if kid == "" {
		return nil, fmt.Errorf("missing jwt kid header claim")
	}

	jwk := getCache().Get(kid)
	if jwk == nil || !getCache().IsValid(kid, clock.Now().UTC()) {
		var err error
		jwk, err = jwt.GetJSONWebKey(ctx, &jwt.GetJSONWebKeyParams{
			KeyID:      kid,
			JWKSClient: jwksClient,
		})
		if err != nil {
			return nil, err
		}
	}
	getCache().Set(kid, jwk, clock.Now().UTC().Add(time.Hour))
	return jwk, nil
}

type AuthorizationParams struct {
	jwt.VerifyParams
	// AuthorizationFailureHandler gets executed when request authorization
	// fails. Pass a custom http.Handler to control the http.Response for
	// invalid authorization. The default is a Response with an empty body
	// and 401 Unauthorized status.
	AuthorizationFailureHandler http.Handler
	// JWKSClient is the jwks.Client that will be used to fetch the
	// JSON Web SecretKey Set. A default client will be used if none is
	// provided.
	JWKSClient *jwks.Client
}

// AuthorizationOption is a functional parameter for configuring
// authorization options.
type AuthorizationOption func(*AuthorizationParams) error

// A cache to store JSON Web Keys.
type jwkCache struct {
	mu      sync.RWMutex
	entries map[string]*cacheEntry
}

// Each entry in the JWK cache has a value and an expiration date.
type cacheEntry struct {
	value     *clerk.JSONWebKey
	expiresAt time.Time
}

// IsValid returns true if a non-expired entry exists in the cache
// for the provided key, false otherwise.
func (c *jwkCache) IsValid(key string, t time.Time) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.entries[key]
	return ok && entry != nil && entry.expiresAt.After(t)
}

// Get fetches the JSON Web SecretKey for the provided key, unless the
// entry has expired.
func (c *jwkCache) Get(key string) *clerk.JSONWebKey {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.entries[key]
	if !ok || entry == nil {
		return nil
	}
	return entry.value
}

// Set stores the JSON Web SecretKey in the provided key and sets the
// expiration date.
func (c *jwkCache) Set(key string, value *clerk.JSONWebKey, expiresAt time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = &cacheEntry{
		value:     value,
		expiresAt: expiresAt,
	}
}

var cacheInit sync.Once

// A "singleton" JWK cache for the package.
var cache *jwkCache

// getCache returns the library's default cache singleton.
// Please note that the returned Cache is a package-level variable.
// Using the package with more than one Clerk API secret keys might
// require to use different Clients with their own Cache.
func getCache() *jwkCache {
	cacheInit.Do(func() {
		cache = &jwkCache{
			entries: map[string]*cacheEntry{},
		}
	})
	return cache
}

func SessionClaimsFromContext(c echo.Context) (*clerk.SessionClaims, bool) {
	value, ok := c.Get(ctxJwtUserClaims).(*clerk.SessionClaims)
	return value, ok
}
