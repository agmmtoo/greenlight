# Greenlight

## Install Dependencies

```bash
go mod download
```

## DB Setup

[migrate](https://github.com/golang-migrate/migrate) is used for db migrations.
Migration files are in `./migrations`

store db dsn as shell var

```bash
export GREENLIGHT_DB_DSN = 'postgres://username:password@localhost/dbname'
```

run migrations

```bash
migrate -path=./migrations -database=$GREENLIGHT_DB_DSN up
```

you can check `schema_migrations` table for migration versions

```bash
psql $GREENLIGHT_DB_DSN
```

```sql
select * from schema_migrations;
```

(For more info, [check migrate docs](https://github.com/golang-migrate/migrate).)

## Start Server

```bash
go run ./cmd/api
```

check available config

```bash
go run ./cmd/api -help
```
