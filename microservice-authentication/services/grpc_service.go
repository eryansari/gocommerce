package services

import (
	"github.com/sirupsen/logrus"
	"gocommerce/microservice-authentication/pkg/pbuff"
	"gocommerce/microservice-authentication/repositories"
	"google.golang.org/grpc"
)

// GRPCService
type GRPCService struct {
	server  *grpc.Server
	AuthSvc *AuthSvc
}

// Register GRPC
func BuildGRPCService(server *grpc.Server, repository repositories.Repository) *GRPCService {
	logger := logrus.StandardLogger()

	authSvc := NewAuthSvc(repository.UserRepository(), logger)
	pbuff.RegisterAuthApiServer(server, authSvc)

	return &GRPCService{
		server:  server,
		AuthSvc: authSvc,
	}
}
