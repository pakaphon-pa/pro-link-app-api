#!/bin/sh
# wait-for-db.sh

set -e
  
host="$1"
shift
  
until nc -z -v -w30 "$host"
do
  >&2 echo "database is unavailable - waiting sql up..."
  sleep 1
done
  
>&2 echo "database is up - executing command"
exec "$@"