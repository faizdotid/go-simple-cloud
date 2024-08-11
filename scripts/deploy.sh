#!/bin/bash

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
  echo "Docker is not installed. Please install Docker and try again."
  exit 1
fi

# Define the migration file path
migration_file="assets/migrations/db.sql"

# Check if the migration file exists
docker exec -i db