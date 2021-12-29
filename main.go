package main

import (
    "log"
    "net"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"

    "github.com/spark8899/go-tag-service/server"
    pb "github.com/spark8899/go-tag-service/proto"
)

func main() {
    port := "8001"
    s := grpc.NewServer()
    pb.RegisterTagServiceServer(s, server.NewTagServer())
    reflection.Register(s)

    lis, err := net.Listen("tcp", ":"+port)
    if err != nil {
        log.Fatalf("net.Listen err: %v", err)
    }

    err = s.Serve(lis)
    if err != nil {
        log.Fatalf("server.Serve err: %v", err)
    }
}
