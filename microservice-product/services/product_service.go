package services

import (
	"context"
	"encoding/hex"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/sirupsen/logrus"
	"gocommerce/microservice-product/pkg/model"
	"gocommerce/microservice-product/pkg/pbuff"
	"gocommerce/microservice-product/repositories"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ProductService
type ProductSvc struct {
	prepo repositories.ProductRepository
	log   *logrus.Logger
}

// NewProductSvc
func NewProductSvc(prepo repositories.ProductRepository, log *logrus.Logger) *ProductSvc {
	return &ProductSvc{
		prepo: prepo,
		log:   log,
	}
}

// CreateProduct
func (s *ProductSvc) Create(ctx context.Context, p *pbuff.ProductDto) (*pbuff.ProductDto, error) {
	var (
		pname = p.Name
		price = p.Price
		stock = p.Stock
		desc  = p.Description
	)

	if pname == "" {
		return nil, status.Errorf(codes.InvalidArgument, "product name required")
	}
	if desc == "" {
		return nil, status.Errorf(codes.InvalidArgument, "description name required")
	}
	if price <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "product price empty")
	}
	if stock <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "product stock required")
	}

	var pdto *pbuff.ProductDto
	rm := &model.Product{
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Stock:       p.Stock,
		Status:      p.Status,
	}

	pres, err := s.prepo.Create(ctx, rm)
	if err != nil {
		s.log.Error(err)
		return nil, status.Errorf(codes.Internal, "Create product failed")
	}

	pid, _ := hex.DecodeString(pres.ID.Hex())
	pdto = &pbuff.ProductDto{
		Id:          string(pid),
		Name:        pres.Name,
		Description: pres.Description,
		Price:       pres.Price,
		Stock:       pres.Stock,
		Status:      pres.Status,
	}

	s.log.Debugf("created Product (ID: %s, Name: %s)", pdto.Id, pdto.Name)

	return pdto, nil

}

// GetProductByID
func (s *ProductSvc) GetProductByID(ctx context.Context, id *wrappers.StringValue) (*pbuff.ProductDto, error) {

	pdetl, err := s.prepo.FindByID(id.GetValue())
	if err != nil {
		s.log.Error(err)
		return nil, status.Errorf(codes.NotFound, "Product not found")
	}

	pid, _ := hex.DecodeString(pdetl.ID.Hex())
	response := &pbuff.ProductDto{
		Id:          string(pid),
		Name:        pdetl.Name,
		Description: pdetl.Description,
		Price:       pdetl.Price,
		Stock:       pdetl.Stock,
		Status:      pdetl.Status,
	}

	return response, nil

}

//ListProducts
func (s *ProductSvc) ListProduct(ctx context.Context, e *empty.Empty) (*pbuff.ProductDtoList, error) {
	var (
		plist *pbuff.ProductDtoList
		pdtos []*pbuff.ProductDto
		pdto  *pbuff.ProductDto
	)

	product, err := s.prepo.FindAll()
	if err != nil {
		s.log.Error(err)
		return nil, status.Errorf(codes.NotFound, "Product not found")
	}

	for _, p := range product {
		pid, _ := hex.DecodeString(p.ID.Hex())
		pdto = &pbuff.ProductDto{
			Id:          string(pid),
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Stock:       p.Stock,
			Status:      p.Status,
		}
		pdtos = append(pdtos, pdto)
	}

	plist = new(pbuff.ProductDtoList)
	plist.List = pdtos

	return plist, nil
}
