package main

import (
    "log"
    "flag"
    "net/http"

    "github.com/grpc-ecosystem/grpc-gateway/runtime"
    pb "github.com/spark8899/go-tag-service/proto"
    "github.com/spark8899/go-tag-service/server"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

var port string

func init() {
    flag.StringVar(&port, "port", "8004", "启动端口号")
    flag.Parse()
}

func main() {
    err := RunServer(port)
    if err != nil {
        log.Fatalf("Run Serve err: %v", err)
    }
}

func RunServer(port string) error {
    httpMux := runHttpServer()
    grpcS := runGrpcServer()
    gatewayMux := runGrpcGatewayServer()

    httpMux.Handle("/", gatewayMux)

    return http.ListenAndServe(":"+port, grpcHandlerFunc(grpcS, httpMux))
}

func runHttpServer() *http.ServeMux {
    serveMux := http.NewServeMux()
    serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
        _, _ = w.Write([]byte(`pong`))
    })

    return serveMux
}

func runGrpcServer() *grpc.Server {
    s := grpc.NewServer()
    pb.RegisterTagServiceServer(s, server.NewTagServer())
    reflection.Register(s)

    return s
}

func runGrpcGatewayServer() *runtime.ServeMux {
    endpoint := "0.0.0.0:" + port
    gwmux := runtime.NewServeMux()
    dopts := []grpc.DialOption{grpc.WithInsecure()}
    _ = pb.RegisterTagServiceHandlerFromEndpoint(context.Background(), gwmux, endpoint, dopts)

    return gwmux
}
