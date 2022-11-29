package backend

import (
	"github.com/RomanShirov/ipblock-service/internal/service"
	"github.com/go-redis/redis/v8"
)

type Server struct {
	service.UnimplementedRPCServiceServer
	Redis *redis.Client
}
