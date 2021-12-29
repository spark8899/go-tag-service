module github.com/spark8899/go-tag-service

go 1.17

require (
	github.com/opentracing/opentracing-go v1.2.0
	google.golang.org/grpc v1.43.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/golang/glog v1.0.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/soheilhy/cmux v0.1.5 // indirect
	golang.org/x/net v0.0.0-20211216030914-fe4d6282115f // indirect
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20211223182754-3ac035c7e7cb // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/spark8899/go-tag-service/pkg/bapi => ./pkg/bapi
	github.com/spark8899/go-tag-service/pkg/errcode => ./pkg/errcode
	github.com/spark8899/go-tag-service/proto => ./proto
)
