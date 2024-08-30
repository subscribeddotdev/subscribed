package http

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echomiddleware "github.com/oapi-codegen/echo-middleware"

	"github.com/subscribeddotdev/subscribed-backend/internal/app/auth"
	"github.com/subscribeddotdev/subscribed-backend/internal/common/logs"
)

func loggerMiddleware(logger *logs.Logger, isDebug bool) echo.MiddlewareFunc {
	if isDebug {
		return middleware.Logger()
	}

	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogStatus:    true,
		LogMethod:    true,
		LogError:     false,
		LogRequestID: true,
		LogHeaders:   []string{echo.HeaderXCorrelationID},
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			lvl := slog.LevelInfo
			msg := "request handled successfully"
			if v.Error != nil || v.Status > http.StatusBadRequest {
				lvl = slog.LevelError
				msg = "request handled with an error"
			}

			logger.LogAttrs(
				c.Request().Context(),
				lvl,
				msg,
				slog.String("method", v.Method),
				slog.String("uri", v.URI),
				slog.Int("status", v.Status),
				slog.String("correlation_id", v.RequestID),
			)

			return nil
		},
	})
}

func errorHandler(logger *logs.Logger) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}

		logger.Debug("Handler failed", "error", err)

		code := http.StatusInternalServerError
		errorSlug := "internal-server-error"
		errorMessage := err.Error()

		switch e := err.(type) {
		case *echo.HTTPError:
			code = e.Code
			errorMessage = e.Error()
		case *HandlerError:
			code = e.Code
			errorSlug = e.Slug()
			errorMessage = e.Error()
		}

		if c.Request().Method == http.MethodHead {
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, ErrorResponse{
				Error:   errorSlug,
				Message: errorMessage,
			})
		}
		if err != nil {
			logger.Debug("Failed to send error response", "error", err)
		}
	}
}

type apiKeyMiddleware struct {
	auth *auth.Service
}

func (a *apiKeyMiddleware) Middleware(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	r := input.RequestValidationInput.Request

	apiKeySecretKey := strings.TrimSpace(r.Header.Get("x-api-key"))
	if apiKeySecretKey == "" {
		return errors.New("x-api-key header cannot be empty")
	}

	apiKey, err := a.auth.ResolveApiKeyFromSecretKey(ctx, apiKeySecretKey)
	if err != nil {
		return err
	}

	eCtx := echomiddleware.GetEchoContext(ctx)
	eCtx.Set("org_id", apiKey.OrgID())
	eCtx.Set("api_key_id", apiKey.Id().String())

	return nil
}
