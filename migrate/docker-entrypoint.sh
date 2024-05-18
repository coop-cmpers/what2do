#!/bin/sh

# Abort on any error
set -e

# Abort if any environment variables are not set
if [ -z "$DB_USER" ] || [ -z "$DB_PASSWORD" ] || [ -z "$DB_NAME" ] || [ -z "$TEST_DB_USER" ] || 
  [ -z "$TEST_DB_PASSWORD" ] || [ -z "$TEST_DB_NAME" ] || [ -z "$PROD_TEST_DB_HOST" ] || 
  [ -z "$PROD_TEST_DB_PORT" ] || [ -z "$PROD_DB_HOST" ] || [ -z "$PROD_DB_PORT" ] ; then
  echo "Failed running docker-entrypoint.sh: one or more environment variables are not set" 1>&2
  exit 1
fi

# TODO: hacky solution
# Wait 15 seconds for Postgres services to be up
echo "Starting 15 seconds wait for Postgres services to be up"
sleep 15s

# Run migrations for prod test DB
# shellcheck disable=SC2086
migrate -path /migrations -database postgres://$TEST_DB_USER:$TEST_DB_PASSWORD@$PROD_TEST_DB_HOST:$PROD_TEST_DB_PORT/$TEST_DB_NAME?sslmode=disable -verbose up

# Run migrations for prod DB
# shellcheck disable=SC2086
migrate -path /migrations -database postgres://$DB_USER:$DB_PASSWORD@$PROD_DB_HOST:$PROD_DB_PORT/$DB_NAME?sslmode=disable -verbose up

exec "$@"