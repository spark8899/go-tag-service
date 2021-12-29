package main

import (
    "context"
    "log"

    pb "github.com/spark8899/go-tag-service/proto"
    "google.golang.org/grpc"
)

func main() {
    ctx := context.Background()
    clientConn, _ := GetClientConn(ctx, "localhost:8001", nil)
    if err != nil {
        log.Fatalf("err: %v", err)
    }
    defer clientConn.Close()
    tagServiceClient := pb.NewTagServiceClient(clientConn)
    resp, err := tagServiceClient.GetTagList(ctx, &pb.GetTagListRequest{Name: "Go"})
    if err != nil {
        log.Fatalf("tagServiceClient.GetTagList err: %v", err)
    }

    log.Printf("resp: %v", resp)
}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
    opts = append(opts, grpc.WithInsecure())
    return grpc.DialContext(ctx, target, opts...)
}
