package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/coop-cmpers/what2do-backend/src/constants"
)

type Store struct {
	db *sql.DB
}

func getConnectionStringFromCtx(ctx context.Context) (string, string, error) {
	env := ctx.Value(constants.Env)
	envMap, ok := env.(map[string]string)
	if !ok {
		log.Fatalf("Failed to cast environment variables in context to map")
		return "", "", errors.New("failed to cast environment variables in context to map")
	}

	// Specify database access mode (prod / prod_test / dev / dev_test)
	mode := envMap["DB_MODE"]
	if len(mode) == 0 {
		log.Fatalf("Failed to find DB access mode")
		return "", "", errors.New("failed to find DB access mode")
	}

	// Specify driver
	driverName := envMap["TEST_DB_DRIVER"]
	if !strings.HasSuffix(mode, "_test") {
		driverName = envMap["DB_DRIVER"]
	}

	// Specify credentials
	connStr := "postgresql://" + envMap["TEST_DB_USER"] + ":" + envMap["TEST_DB_PASSWORD"] + "@"
	if !strings.HasSuffix(mode, "_test") {
		connStr = "postgresql://" + envMap["DB_USER"] + ":" + envMap["DB_PASSWORD"] + "@"
	}

	// Specify address
	switch mode {
	case "prod":
		connStr += envMap["PROD_DB_HOST"] + ":" + envMap["PROD_DB_PORT"]
	case "prod_test":
		connStr += envMap["PROD_TEST_DB_HOST"] + ":" + envMap["PROD_TEST_DB_PORT"]
	case "dev":
		connStr += envMap["DEV_DB_HOST"] + ":" + envMap["DEV_DB_PORT"]
	case "dev_test":
		connStr += envMap["DEV_TEST_DB_HOST"] + ":" + envMap["DEV_TEST_DB_PORT"]
	default:
		log.Fatalf("Failed because of an unexpected DB access mode")
		return "", "", fmt.Errorf("failed because of an unexpected DB access mode - mode: %s", mode)
	}

	// Specify database
	if strings.HasSuffix(mode, "_test") {
		connStr += "/" + envMap["TEST_DB_NAME"]
	} else {
		connStr += "/" + envMap["DB_NAME"]
	}

	// Disable SSL
	connStr += "?sslmode=disable"

	return driverName, connStr, nil
}

func (s *Store) Connect(ctx context.Context) (*sql.DB, error) {
	driverName, connStr, err := getConnectionStringFromCtx(ctx)
	if err != nil {
		return nil, err
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

	s.db = db
	return db, nil
}
