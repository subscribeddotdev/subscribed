package postgres

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/pressly/goose"
)

func Connect(uri string) (*sql.DB, error) {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, fmt.Errorf("could not open a connection to postgres: \nuri: %s\nerror: %v", uri, err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("unable to ping postgres: \nuri: %s\nerror: %v", uri, err)
	}

	return db, nil
}

func ApplyMigrations(db *sql.DB, dir string) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	workdir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("unable to get the current working directory: %v", err)
	}

	if err = goose.Up(db, fmt.Sprintf("%s/%s", workdir, dir)); err != nil {
		return fmt.Errorf("unable to apply migrations: %v", err)
	}

	return nil
}
