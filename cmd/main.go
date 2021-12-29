package main

import (
    "flag"
    "log"
    "net"

    pb "github.com/spark8899/go-tag-service/proto"
    "github.com/spark8899/go-tag-service/server"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

var grpcPort string
var httpPort string

func init() {
    flag.StringVar(&grpcPort, "grpc_port", "8004", "gRPC启动端口号")
    flag.Parse()
}

func main() {
    errs := make(chan error)
    go func() {
        err := RunGrpcServer(grpcPort)
        if err != nil {
            errs <- err
        }
    }()

    select {
    case err := <-errs:
        log.Fatalf("Run Server err: %v", err)
    }
}

func RunGrpcServer(port string) error {
    s := grpc.NewServer()
    pb.RegisterTagServiceServer(s, server.NewTagServer())
    reflection.Register(s)
    lis, err := net.Listen("tcp", ":"+port)
    if err != nil {
        return err
    }

    err = s.Serve(lis)
    if err != nil {
        return err
    }

    return nil
}
