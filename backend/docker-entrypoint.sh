#!/bin/sh

# Abort on any error
set -e

# Abort if any environment variables are not set
if [ -z "$PROD_DB_HOST" ] || [ -z "$PROD_DB_PORT" ] || [ -z "$PROD_TEST_DB_HOST" ] || [ -z "$PROD_TEST_DB_PORT" ]; then
  echo "Failed running docker-entrypoint.sh: one or more environment variables are not set" 1>&2
  exit 1
fi

# Check that prod DB is up
./wait-for-it.sh "$PROD_DB_HOST:$PROD_DB_PORT"

# Check that prod test DB is up
./wait-for-it.sh "$PROD_TEST_DB_HOST:$PROD_TEST_DB_PORT"

# Run the main container command
exec "$@"
