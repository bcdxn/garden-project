# Database Migrations

## Prerequisites

Use `golang-migrations` tool to complete migrations

```sh
brew install golang-migrate 
```

Use official Postgres Docker image: https://hub.docker.com/_/postgres

```sh
docker pull postgres:latest
```

## Creating a Database From Scratch (For Local Dev)

### Create Database

```sh
docker run \
  -ti \
  --rm \
  -e POSTGRES_HOST_AUTH_METHOD=trust \
  -p 5432:5432 postgres:latest
```

### Run Migrations (Up)

```sh
migrate  \
  -database "postgresql://postgres@0.0.0.0:5432/postgres?sslmode=disable" \
  -path db/migrations up
```
