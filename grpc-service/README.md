# GPRC Service

Opinionated [GRPC](https://grpc.io/) service with embedded [GRPC Gateway](https://github.com/grpc-ecosystem/grpc-gateway) for REST HTTP.

- grpc code generated with [Buf CLI](https://buf.build/docs/installation)
- postgres for SQL database
- sql migrations with [Goose](https://github.com/pressly/goose)
- generated sql methods with [sqlc](https://sqlc.dev/) with [pgx](https://github.com/jackc/pgx)

## Install required tools

```shell
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc \
    github.com/sqlc-dev/sqlc/cmd/sqlc \
    github.com/pressly/goose/v3/cmd/goose
```

Install [Buf CLI](https://buf.build/docs/installation)

## Generate gprc code and swagger json

```shell
buf generate
```

## Generate db methods from sql queries

```shell
sqlc generate
```

## Run service locally
```shell
go run . serve
```

## Adding SQL migrations

```shell
goose -dir db/migrations create DESCRIPTIVE_FILENAME sql
```

## [GRPC UI](https://github.com/fullstorydev/grpcui) to test GRPC endpoints (like a Postman for HTTP)