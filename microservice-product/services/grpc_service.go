package services

import (
	"github.com/sirupsen/logrus"
	"gocommerce/microservice-product/pkg/pbuff"
	"gocommerce/microservice-product/repositories"
	"google.golang.org/grpc"
)

// GRPCService ...
type GRPCService struct {
	server     *grpc.Server
	ProductSvc *ProductSvc
}

// BuildGRPCService register gRPC services implementation
func BuildGRPCService(server *grpc.Server, repository repositories.Repository) *GRPCService {
	logger := logrus.StandardLogger()

	productSvc := NewProductSvc(repository.ProductRepository(), logger)
	pbuff.RegisterProductApiServer(server, productSvc)

	return &GRPCService{
		server:     server,
		ProductSvc: productSvc,
	}
}
