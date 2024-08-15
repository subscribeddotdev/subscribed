package http

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/oapi-codegen/echo-middleware"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
)

type jwtCustomClaims struct {
	jwt.RegisteredClaims
	OrganizationID string `json:"organization_id"`
	MemberID       string `json:"member_id"`
	Email          string `json:"email"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
}

func signJwt(member *iam.Member, secret string) (string, error) {
	secret = strings.TrimSpace(secret)
	if secret == "" {
		return "", errors.New("secret cannot be empty")
	}

	claims := jwtCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "subscriber-backend",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		OrganizationID: member.OrgID().String(),
		MemberID:       member.ID().String(),
		Email:          member.Email().String(),
		FirstName:      member.FirstName(),
		LastName:       member.LastName(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("error signing token: %v", err)
	}

	return signedToken, nil
}

func decodeJwt(signedToken, secret string) (*jwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(signedToken, &jwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("error decoding jwt: %v", err)
	}

	claims, ok := token.Claims.(*jwtCustomClaims)
	if !ok {
		return nil, fmt.Errorf("error parsing the jwt claims: %v", err)
	}

	return claims, nil
}

type jwtMiddleware struct {
	secret string
}

func (j *jwtMiddleware) Middleware(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	r := input.RequestValidationInput.Request

	signedToken := strings.TrimPrefix(strings.TrimSpace(r.Header.Get(echo.HeaderAuthorization)), "Bearer ")
	if signedToken == "" {
		return errors.New("jwt signedToken cannot be empty")
	}

	claims, err := decodeJwt(signedToken, j.secret)
	if err != nil {
		return NewHandlerErrorWithStatus(err, "error-validating-jwt", http.StatusUnauthorized)
	}

	eCtx := echomiddleware.GetEchoContext(ctx)
	eCtx.Set("user_claims", claims)

	return nil
}
