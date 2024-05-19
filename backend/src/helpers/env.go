package helpers

import (
	"context"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/coop-cmpers/what2do-backend/src/constants"
	"github.com/joho/godotenv"
)

// Return map from OS environment variables
func getEnvFromOS() map[string]string {
	envMap := make(map[string]string)

	for _, envPair := range os.Environ() {
		if i := strings.Index(envPair, "="); i >= 0 {
			envKey := envPair[:i]
			envValue := envPair[i+1:]
			envMap[envKey] = envValue
		}
	}

	return envMap
}

// Add environment variables to context
func AddEnvToCtx(ctx context.Context) context.Context {
	// Running locally: will read environment variables from .env file
	env, err := godotenv.Read("../.env")
	if err != nil {
		// Running in Docker: will read environment variables from OS
		log.Printf("Could not read local .env file - assuming all environment variables are in the OS - err: %v", err)
		env = getEnvFromOS()
	}

	ctx = context.WithValue(ctx, constants.Env, env)
	return ctx
}

func GetEnvFromCtx(ctx context.Context) (map[string]string, error) {
	env := ctx.Value(constants.Env)

	envMap, ok := env.(map[string]string)
	if !ok {
		log.Fatalf("Failed to cast environment variables in context to map")
		return nil, errors.New("failed to cast environment variables in context to map")
	}

	return envMap, nil
}
