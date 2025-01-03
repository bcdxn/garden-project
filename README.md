# Garden Project

_Plan and track your garden_

## Developer

### Building

`todo`

### Testing

`todo`

### Running Locally

#### Database

##### Start Postgres

```sh
docker run \
  -ti \
  --rm \
  -e POSTGRES_HOST_AUTH_METHOD=trust \
  -p 5432:5432 postgres:latest
```

##### Run Migrations

```sh
DB_URI="postgresql://postgres@0.0.0.0:5432/postgres?sslmode=disable" \
go run cmd/db/migrations/main.go up
```

##### Seed Data

```sh
DB_URI="postgresql://postgres@0.0.0.0:5432/postgres?sslmode=disable" \
go run cmd/db/seeds/main.go up
```

#### Start API

```sh
DB_URI="postgresql://postgres@0.0.0.0:5432/postgres?sslmode=disable" \
GP_PORT=8080 \
go run cmd/rest_api/main.go
```

#### View API Documentation

OpenAPI Spec is available via
- hosted SwaggerUI: http://localhost8080/docs/api/v1
- raw format: 
