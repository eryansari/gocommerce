package services

import (
	"context"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gocommerce/microservice-product/pkg/model"
	"gocommerce/microservice-product/pkg/pbuff"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (r ProductRepositoryMock) Create(context.Context, *model.Product) (model.Product, error) {
	args := r.Called()
	products := model.Product{
		Name: "tea",
	}
	return products, args.Error(1)
}

func (r ProductRepositoryMock) FindByID(string) (model.Product, error) {
	args := r.Called()
	products := model.Product{
		ID: "62c4f9e3e01960e5391bd627",
	}
	return products, args.Error(1)
}

func (r ProductRepositoryMock) FindAll() ([]model.Product, error) {
	args := r.Called()
	products := []model.Product{
		{Name: "water", Description: "water is good", Price: 111, Stock: 5},
		{Name: "tea", Description: "tea is good", Price: 111, Stock: 5},
	}
	return products, args.Error(1)
}

func TestService_Create(t *testing.T) {
	repositories := ProductRepositoryMock{}
	repositories.On("Create").Return(model.Product{}, nil)

	ctx := context.Background()

	service := ProductSvc{repositories, logrus.StandardLogger()}
	req := &pbuff.ProductDto{
		Name:        "tea",
		Description: "tea is good",
		Price:       111,
		Stock:       5,
	}
	product, _ := service.Create(ctx, req)
	assert.Equal(t, product.Name, "tea")
}

func TestService_GetProductByID(t *testing.T) {
	repositories := ProductRepositoryMock{}
	repositories.On("FindByID").Return(model.Product{}, nil)

	ctx := context.Background()

	service := ProductSvc{repositories, logrus.StandardLogger()}
	product, _ := service.GetProductByID(ctx, &wrapperspb.StringValue{})
	assert.Equal(t, product.Id, "62c4f9e3e01960e5391bd627")
}

func TestService_ListProduct(t *testing.T) {
	repositories := ProductRepositoryMock{}
	repositories.On("FindAll").Return(model.Product{}, nil)

	ctx := context.Background()

	service := ProductSvc{repositories, logrus.StandardLogger()}
	product, _ := service.ListProduct(ctx, &emptypb.Empty{})
	assert.NotEmpty(t, product)
}
