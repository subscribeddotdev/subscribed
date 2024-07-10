package components_test

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/common/logs"
	"github.com/subscribeddotdev/subscribed-backend/internal/common/postgres"
	"github.com/subscribeddotdev/subscribed-backend/misc/tools/wait/wait_for"
	"github.com/subscribeddotdev/subscribed-backend/tests/client"
)

var (
	db  *sql.DB
	ctx context.Context
)

func TestMain(m *testing.M) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Minute*2)
	defer cancel()

	// Wait for postgres and rabbitmq and other default dependencies running in containers
	wait_for.Run()

	// Wait for the backend to be spun up
	waitfor := wait_for.NewWaitFor(logs.New())
	waitfor.Do(func() error {
		req, err := http.Get(fmt.Sprintf("http://localhost:%s/health", os.Getenv("HTTP_PORT")))
		if err != nil {
			return err
		}

		if req.StatusCode != http.StatusOK {
			return fmt.Errorf("expected status code 200 instead got %d", req.StatusCode)
		}

		return nil
	}, "subscribed-backend", time.Second*30)
	waitfor.Wait()

	var err error
	db, err = postgres.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	err = postgres.ApplyMigrations(db, "../../misc/sql/migrations")
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func getClient(t *testing.T, token string) *client.ClientWithResponses {
	cli, err := client.NewClientWithResponses(
		fmt.Sprintf("http://localhost:%s", os.Getenv("HTTP_PORT")),
		client.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			if token != "" {
				req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
			}
			return nil
		}),
	)
	require.NoError(t, err)

	return cli
}

func getClientWithApiKey(t *testing.T, key string) *client.ClientWithResponses {
	cli, err := client.NewClientWithResponses(
		fmt.Sprintf("http://localhost:%s", os.Getenv("HTTP_PORT")),
		client.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			if key == "" {
				return errors.New("api key is missing")
			}

			req.Header.Set("x-api-key", key)

			return nil
		}),
	)
	require.NoError(t, err)

	return cli
}

func toPtr[T any](v T) *T {
	return &v
}
