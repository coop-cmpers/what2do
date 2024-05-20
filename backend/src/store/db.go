package store

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	pb "github.com/coop-cmpers/what2do-backend/protos-gen/helloworld/v1"
	"github.com/coop-cmpers/what2do-backend/src/constants"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type StoreInterface interface {
	Connect(ctx context.Context) error
	InitialiseLogger(ctx context.Context) error
	HelloDB(ctx context.Context, id int32) (*pb.TestDBObject, error)
	CreateEvent(ctx context.Context, event *Event) error
	GetEvent(ctx context.Context, eventID string) (*Event, error)
}

type Store struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

func (s *Store) getConnectionStringFromCtx(ctx context.Context) (string, string, error) {
	env := ctx.Value(constants.Env)
	envMap, ok := env.(map[string]string)
	if !ok {
		s.logger.Errorf("Failed to cast environment variables in context to map")
		return "", "", errors.New("failed to cast environment variables in context to map")
	}

	// Specify database access mode (prod / prod_test / dev / dev_test)
	mode := envMap["DB_MODE"]
	if len(mode) == 0 {
		s.logger.Errorf("Failed to find DB access mode")
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
		s.logger.Errorf("Failed because of an unexpected DB access mode")
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

func (s *Store) SetLogger(ctx context.Context) error {
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	defer logger.Sync()

	sugaredLogger := logger.Sugar()
	s.logger = sugaredLogger

	return nil
}

func (s *Store) SetDB(ctx context.Context) error {
	driverName, connStr, err := s.getConnectionStringFromCtx(ctx)
	if err != nil {
		return err
	}

	db, err := sqlx.Open(driverName, connStr)
	if err != nil {
		s.logger.Errorf("Failed to connect to postgres database - err: %v", err)
		return err
	}

	err = db.Ping()
	if err != nil {
		s.logger.Errorf("Failed to ping database - err: %v", err)
		return err
	}

	s.db = db
	return nil
}

func NewStore(ctx context.Context) (*Store, error) {
	store := &Store{}

	err := store.SetLogger(ctx)
	if err != nil {
		log.Printf("Failed to initialise database Zap logger - err: %v", err)
		return nil, err
	}

	err = store.SetDB(ctx)
	if err != nil {
		store.logger.Errorf("Failed to initialise databsae connection - err: %v", err)
		return nil, err
	}

	return store, nil
}
