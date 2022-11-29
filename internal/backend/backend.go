package backend

import (
	"context"
	"github.com/RomanShirov/ipblock-service/internal/service"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type Server struct {
	service.UnimplementedRPCServiceServer
	Redis *redis.Client
}

func (s *Server) BlockIP(ctx context.Context, in *service.BlockIPRequest) (*service.BlockIPResponse, error) {
	expirationTime := time.Duration(in.Duration) * time.Second
	err := s.Redis.Set(ctx, in.Ip, "", expirationTime).Err()
	if err != nil {
		log.Fatalf("Redis write error: %v", err)
	}
	return &service.BlockIPResponse{IsBlocked: true}, nil
}

func (s *Server) IsBlockedIP(ctx context.Context, in *service.IsBlockedIPRequest) (*service.IsBlockedIPResponse, error) {
	var isBlocked bool

	_, err := s.Redis.Get(ctx, in.Ip).Result()
	if err == redis.Nil {
		isBlocked = false
	} else if err != nil {
		isBlocked = false
		log.Fatalf("Redis get error: %v", err)
	} else {
		isBlocked = true
	}

	return &service.IsBlockedIPResponse{IsBlocked: isBlocked}, nil
}

func (s *Server) UnblockIP(ctx context.Context, in *service.UnblockIPRequest) (*service.UnblockIPResponse, error) {
	isIpRemoved := s.Redis.Del(ctx, in.Ip)
	if isIpRemoved == nil {
		log.Fatalf("Redis delete error")
	}

	return &service.UnblockIPResponse{IsBlocked: false}, nil
}
