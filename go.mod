module github.com/spark8899/go-tag-service

go 1.17

require (
	github.com/opentracing/opentracing-go v1.2.0
	google.golang.org/grpc v1.43.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/golang/protobuf v1.5.0 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd // indirect
	golang.org/x/text v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
)

replace (
	github.com/spark8899/go-tag-service/pkg/bapi => ./pkg/bapi
	github.com/spark8899/go-tag-service/pkg/errcode => ./pkg/errcode
	github.com/spark8899/go-tag-service/proto => ./proto
)
