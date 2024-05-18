## what2do-migrate

Migration scripts for the Postgres DB, powered by golang-migrate.

Create empty migration scripts (inside ./migrate): ` migrate create -ext sql -dir ./migrations -seq -digits 3 <migration_name>`

To migrate: 
1. Make sure the Postgres services are running
2. Run the `migrate` container