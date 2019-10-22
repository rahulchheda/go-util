package main

import (
	//"context"
	"fmt"
	"github.com/litmuschaos/go-util/internal/grpc"
	//"github.com/litmuschaos/go-util/internal/grpc/domain"
	"github.com/litmuschaos/go-util/internal/grpc/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func getNetListerner(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))

	}

	return lis
}

func main() {
	netListener := getNetListerner(7000)
	opts := []grpc.ServerOption{}
	gRPCServer := grpc.NewServer(opts...)
	executor := impl.NewExecuteFileServiceGrpc()
	service.RegisterExecutorServiceServer(gRPCServer, executor)
	if err := gRPCServer.Serve(netListener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	fmt.Println("gRPC In Action!")
}
