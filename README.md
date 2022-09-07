# MICROSERVICES TRAINING, GO, GRPC

## Proto Gen

```shell
protoc -I=proto proto/*.proto --go_out=:pb --go-grpc_out=:pb

protoc --go-grpc_out=. proto/*.proto
```

## Test with grpcui

```
grpcui -plaintext 127.0.0.1:2222
```

![project-structure.png](./project-structure.png)

go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u google.golang.org/protobuf/{proto,protoc-gen-go}