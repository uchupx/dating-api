ENV_FILE="../.env"
echo "Loading environment variables from $ENV_FILE"
export $(grep -E '^DATABASE' $ENV_FILE | xargs)

## check dependency
if ! command -v migrate &>/dev/null; then
  echo "golang migrate could not be found, please install it to proceed."
  echo "to install follow this link: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate"
  exit 1
fi

if [ $1 == "create" ]; then
  migrate create -ext sql -dir ../migration $2
  exit 0
fi

if [ $1 == "fix-version" ]; then
  cat current_version | xargs migrate -path ../migration -database "mysql://$DATABASE_USERNAME:@tcp($DATABASE_HOST:$DATABASE_PORT)/$DATABASE_NAME" force
  exit 0
fi

migrate -path ../migration -database "mysql://$DATABASE_USERNAME:@tcp($DATABASE_HOST:$DATABASE_PORT)/$DATABASE_NAME" $@

if [ $? -ne 0 ]; then
  echo "Migration failed"
  exit 1
else
  echo $(migrate -path ../migration -database "mysql://$DATABASE_USERNAME:@tcp($DATABASE_HOST:$DATABASE_PORT)/$DATABASE_NAME" version) > current_version
  echo "Migration success"
fi
