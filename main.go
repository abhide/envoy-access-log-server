package main

import (
	"log"
	"net"

	pb "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v3"
	"google.golang.org/grpc"
)

type ALSServer struct {
}

func (a *ALSServer) StreamAccessLogs(logStream pb.AccessLogService_StreamAccessLogsServer) error {
	log.Println("Streaming access logs")
	for {
		data, err := logStream.Recv()
		if err != nil {
			return err
		}
		log.Printf("Received log data: %s\n", data.String())
	}
}

func NewALSServer() *ALSServer {
	return &ALSServer{}
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Failed to start listener on port 8080: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterAccessLogServiceServer(grpcServer, NewALSServer())
	log.Println("Starting ALS Server")
	grpcServer.Serve(listener)
}
