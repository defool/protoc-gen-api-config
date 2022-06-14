# protoc-gen-api-config

`protoc-gen-api-config` is a protoc plugin to make standard URL path based on grpc-gateway [gRPC API Configuration](https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/grpc_api_configuration/).

[中文介绍](./README.zh.md)

## Install

```
go install github.com/defool/protoc-gen-api-config
```

## Example
 
 ./example/foo/v1/api.proto: 
```
syntax = "proto3";

package foo.v1;
option go_package="foo/v1";


message SayHelloRequest {
  string name = 1;
}

message SayHelloResponse {
  string reply = 1;
}

service ExampleService {
  rpc SayHello (SayHelloRequest) returns (SayHelloResponse);
}
```

generate stub file in two step by protoc:
```
protoc  -I . --go_out="./example/generated"  ./example/foo/v1/api.proto
protoc  -I . --api-config_out="./example/generated"  ./example/foo/v1/api.proto
```

OR `buf` use case:

buf.gen.yaml:
```
version: v1
plugins:
  - name: go
    out: generated
    opt: paths=source_relative
  - name: go-grpc
    out: generated
    opt:
      - paths=source_relative
  - name: api-config
    out: generated
    opt: paths=source_relative
```

buf.gen.second.yaml:
```
version: v1
plugins:  
  - name: grpc-gateway
    out: generated
    opt:
    - grpc_api_configuration=generated/api-config.yaml
    - paths=source_relative
  - name: openapiv2
    out: generated
    opt:
    - grpc_api_configuration=generated/api-config.yaml
```

generate stub file in two step：
```
buf generate
buf generate --template buf.gen.second.yaml
```