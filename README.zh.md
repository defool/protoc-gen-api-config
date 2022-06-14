# protoc-gen-api-config

`protoc-gen-api-config`是基于grpc-gateway的[gRPC API Configuration](https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/grpc_api_configuration/)能力生成标准URL路径的 `protoc` 插件。

[English](./README.md)

## 安装

```
go install github.com/defool/protoc-gen-api-config
```

## 示例
 
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

使用protc生成桩代码:
```
protoc  -I . --go_out="./example/generated"  ./example/foo/v1/api.proto
protoc  -I . --api-config_out="./example/generated"  ./example/foo/v1/api.proto
```

或使用`buf`来生成

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

需要执行二步来生成桩代码：
```
buf generate
buf generate --template buf.gen.second.yaml
```