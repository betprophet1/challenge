#!/bin/bash

set -e
set -u

function create_user_and_database_with_postgis() {
  local database=$1
  echo "  Creating user and database '$database'"
  psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    CREATE USER $database;
    CREATE DATABASE $database;
EOSQL
}
