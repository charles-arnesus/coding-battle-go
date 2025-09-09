#!/bin/bash
set -e

# Wait until Postgres is ready
echo "Waiting for Postgres..."
until pg_isready -h postgres -U myuser; do
  sleep 1
done

echo "Postgres is ready! Starting CLI..."
./main
