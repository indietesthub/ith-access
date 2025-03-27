#!/bin/bash

# Load environment variables from .env file
set -a
source /app/.env
set +a

# Run the application
exec /app/api 