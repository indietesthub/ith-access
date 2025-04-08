#!/bin/sh
echo "Waiting for PostgreSQL to be ready..."

until pg_isready -h "$DB_HOST" -p 5432 -U "$DB_USER"; do
  sleep 1
done

echo "PostgreSQL is ready!"
# Load environment variables from .env file
set -a
source /app/.env
set +a

# Run the application
exec /app/api 