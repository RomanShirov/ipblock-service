package main

import (
	"github.com/RomanShirov/ipblock-service/internal/backend"
	"github.com/RomanShirov/ipblock-service/internal/service"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	rdb := redis.NewClient(backend.InitRedisConfig())

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	service.RegisterRPCServiceServer(grpcServer, &backend.Server{
		Redis: rdb,
	})
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Starting gRPC service error: %v", err)
	}
}
