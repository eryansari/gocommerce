package services

import (
	"github.com/sirupsen/logrus"
	"gocommerce/microservice-order/pkg/pbuff"
	"gocommerce/microservice-order/repositories"
	"google.golang.org/grpc"
)

// GRPCService ...
type GRPCService struct {
	server   *grpc.Server
	OrderSvc *OrderSvc
}

// BuildGRPCService register gRPC services implementation
func BuildGRPCService(server *grpc.Server, repository repositories.Repository) *GRPCService {
	logger := logrus.StandardLogger()

	orderSvc := NewOrderSvc(repository.OrderRepository(), logger)
	pbuff.RegisterOrderApiServer(server, orderSvc)

	return &GRPCService{
		server:   server,
		OrderSvc: orderSvc,
	}
}
