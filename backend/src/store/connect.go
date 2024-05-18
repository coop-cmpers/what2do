package store

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/coop-cmpers/what2do-backend/src/constants"
)

func Connect(ctx context.Context) (*sql.DB, error) {
	env := ctx.Value(constants.CtxEnv)
	envMap, ok := env.(map[string]string)
	if !ok {
		log.Fatalf("Failed to cast environment variables in context to map")
		return nil, errors.New("failed to cast environment variables in context to map")
	}

	// Default to using test db
	driverName := envMap["TEST_DB_DRIVER"]
	connStr := "postgresql://" + envMap["TEST_DB_USER"] + ":" + envMap["TEST_DB_PASSWORD"] + "@" + envMap["TEST_DB_HOST"] + "/" + envMap["TEST_DB_NAME"] + "?sslmode=disable"

	// If mode is `prod`, use prod db
	if envMap["mode"] == "prod" {
		driverName = envMap["DB_DRIVER"]
		connStr = "postgresql://" + envMap["DB_USER"] + ":" + envMap["DB_PASSWORD"] + "@" + envMap["DB_HOST"] + "/" + envMap["DB_NAME"] + "?sslmode=disable"
	}

	db, err := sql.Open(driverName, connStr)
	if err != nil {
		log.Fatalf("Failed to connect to postgres database - err: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database - err: %v", err)
		return nil, err
	}

	return db, nil
}
