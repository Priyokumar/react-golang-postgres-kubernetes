## Postgres docker run commmand

docker run -d -p 5432:5432 --name postgres -e POSTGRES_USER=test -e POSTGRES_PASSWORD=test -e POSTGRES_DB=test -e PGDATA=/var/lib/postgresql/data/pgdata -v /Users/laxmihuidrom/pg/data:/var/lib/postgresql/data postgres:16.0

## Docker Build MAC

docker buildx  build --platform=linux/amd64  --no-cache --progress plain -t us-central1-docker.pkg.dev/brilliant-dock-408207/priyo/go-react-postgres/frontend .




