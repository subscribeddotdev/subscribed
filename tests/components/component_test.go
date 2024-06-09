package components_test

import (
	"context"
	"database/sql"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/subscribeddotdev/subscribed-backend/internal/common/postgres"
)

var (
	db  *sql.DB
	ctx context.Context
)

func TestMain(m *testing.M) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Minute*2)
	defer cancel()

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

/*func getClient(t *testing.T) *client.ClientWithResponses {
	cli, err := client.NewClientWithResponses(
		fmt.Sprintf("http://localhost:%s", os.Getenv("HTTP_PORT")),
		client.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			return nil
		}),
	)
	require.NoError(t, err)

	return cli
}
*/
