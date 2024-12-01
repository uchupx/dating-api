#!/bin/sh

ENV_FILE="../.env"
echo "Loading environment variables from $ENV_FILE"
export $(grep -Ev '^RSA|^##' $ENV_FILE | xargs)

docker-compose -f ../docker-compose.yml $@
