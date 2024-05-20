# What2Do

To run: `docker compose -f "docker-compose.yml" up -d --build`

This project is a monorepo consisting of five directories:
  - `/frontend` stores the React / Typescript (+ Material UI) frontend service
  - `/backend` stores the Go backend service
  - `/envoy` stores the Envoy proxy service
  - `/protobufs` stores the Protocol Buffers, which are compiled into the below directories:
    - `/frontend/src/protos-gen` for TypeScript and Connect-Web
    - `/backend/protos-gen` for Go and gRPC
  - `/postgres` stores the Postgres schema and data for a Production and Test server
  - `/migrate` stores the migration scripts for the dabase, using `golang-migrate`

There are six services, currently built using Docker and Docker Compose:
  - `what2do-frontend` listening on port 3000:3000 (left is external, right is internal)
  - `what2do-backend` listening on port 12345:12345
  - `what2do-envoy` listening on port 10000:10000
  - `what2do-postgres` listening on port 20000:5432
  - `what2do-postgres-test` listening on port 20001:5432
  - `what2do-migrate` for database migrations