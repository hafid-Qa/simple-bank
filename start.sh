#!/bin/sh

set -e
echo "migrating the database"
/app/migrate -path db/migrations -database "$DB_SOURCE" -verbose up
echo "start the app"
exec "$@"