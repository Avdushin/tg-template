#!/bin/sh
# wait-for-db.sh
set -e

host="$1"
shift
until pg_isready -h "$host" -U "tguser"; do
  echo "Waiting for postgres at $host..."
  sleep 2
done

exec "$@"
