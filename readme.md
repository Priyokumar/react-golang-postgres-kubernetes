## Postgres docker run commmand

docker run -d -p 5432:5432 --name postgres -e POSTGRES_USER=test -e POSTGRES_PASSWORD=test -e POSTGRES_DB=test -e PGDATA=/var/lib/postgresql/data/pgdata -v /Users/laxmihuidrom/pg/data:/var/lib/postgresql/data postgres:16.0


