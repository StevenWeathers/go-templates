# GPRC Service

[GRPC](https://grpc.io/) service with embedded [GRPC Gateway](https://github.com/grpc-ecosystem/grpc-gateway) for REST HTTP

## Install GPRC tools

```shell
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

Install [Buf CLI](https://buf.build/docs/installation)

## Generate gprc code and swagger json

```shell
buf generate
```

## Run service locally
```shell
go run . serve
```

## [GRPC UI](https://github.com/fullstorydev/grpcui) to test GRPC endpoints (like a Postman for HTTP) 