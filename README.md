# go-tag-service
tech go gRPC

# create proto
```
protoc --go_out=plugins=grpc:. ./proto/*.proto
```

# grpcurl test
```
$ grpcurl -plaintext localhost:8001 list
grpc.reflection.v1alpha.ServerReflection
proto.TagService

$ grpcurl -plaintext localhost:8001 list proto.TagService
proto.TagService.GetTagList

$ grpcurl -plaintext -d '{"name":"11"}' localhost:8001 proto.TagService.GetTagList  
{
  "list": [
    {
      "id": "1",
      "name": "11",
      "state": 1
    }
  ],
  "pager": {
    "page": "1",
    "pageSize": "10",
    "totalRows": "1"
  }
}
```

# cmd/362 http and gRPC
```
$ grpcurl -plaintext localhost:8001 proto.TagService.GetTagList

$ curl http://127.0.0.1:8002/ping
```

# cmd/363 http and gRPC
```
$ grpcurl -plaintext localhost:8003 proto.TagService.GetTagList

$ curl http://127.0.0.1:8003/ping
```

# download annotations.proto
```
mkdir -p $GOPATH/src/google/api
cd $GOPATH/src/google/api
wget https://github.com/googleapis/googleapis/raw/master/google/api/annotations.proto

protoc --proto_path=. --proto_path=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=. --go-grpc_out=. --grpc-gateway_out=. ./proto/*.proto
```

