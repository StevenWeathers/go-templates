# GPRC Service

[GRPC](https://grpc.io/) service with embedded [GRPC Gateway](https://github.com/grpc-ecosystem/grpc-gateway) for REST HTTP

Recommended UI Tool to test GRPC endpoints: [GRPC UI](https://github.com/fullstorydev/grpcui)

## Install GPRC tools

```shell
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

Install [Buf CLI](https://buf.build/docs/installation)