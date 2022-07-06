package services

import (
	"context"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gocommerce/microservice-order/pkg/model"
	"gocommerce/microservice-order/pkg/pbuff"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type OrderRepositoryMock struct {
	mock.Mock
}

func (r OrderRepositoryMock) Create(context.Context, *model.Order) (model.Order, error) {
	args := r.Called()
	order := model.Order{
		NoInvoice: "INV/111/111",
		UserID:    "111",
		OrderDetail: []model.OrderDetail{
			{
				ProductID: "62c4f9e3e01960e5391bd627",
				Qty:       5,
			},
		},
	}
	return order, args.Error(1)
}

func (r OrderRepositoryMock) FindByID(string) (model.Order, error) {
	args := r.Called()
	order := model.Order{
		ID: "111",
	}
	return order, args.Error(1)
}

func (r OrderRepositoryMock) FindByUserID(string) ([]model.Order, error) {
	args := r.Called()
	order := []model.Order{
		{NoInvoice: "INV/111/111", UserID: "111", Status: false, OrderDetail: []model.OrderDetail{{ProductID: "62c4f9e3e01960e5391bd627", Qty: 5}}},
		{NoInvoice: "INV/222/222", UserID: "111", Status: false, OrderDetail: []model.OrderDetail{{ProductID: "62c4f9e3e01960e5391bd627", Qty: 2}}},
	}
	return order, args.Error(1)
}

func TestService_Create(t *testing.T) {
	repository := OrderRepositoryMock{}
	repository.On("Create").Return(model.Order{}, nil)

	ctx := context.Background()

	service := OrderSvc{repository, logrus.StandardLogger()}
	orderReq := []*pbuff.OrderDetailDto{{ProductId: "62c4f9e3e01960e5391bd627", Qty: 5}}
	req := &pbuff.OrderReqDto{
		Order: orderReq,
	}
	order, _ := service.Create(ctx, req)
	assert.Equal(t, order.NoInv, "INV/111/111")
	assert.Equal(t, order.UserId, "111")
}

func TestService_GetOrderByID(t *testing.T) {
	repository := OrderRepositoryMock{}
	repository.On("FindByID").Return(model.Order{}, nil)

	ctx := context.Background()

	service := OrderSvc{repository, logrus.StandardLogger()}
	order, _ := service.GetOrderByID(ctx, &wrapperspb.StringValue{})
	assert.Equal(t, order.Id, "111")
}

func TestService_ListOrder(t *testing.T) {
	repository := OrderRepositoryMock{}
	repository.On("FindByUserID").Return(model.Order{}, nil)

	ctx := context.Background()

	service := OrderSvc{repository, logrus.StandardLogger()}
	order, _ := service.ListOrderByUserID(ctx, &wrapperspb.StringValue{})
	assert.NotEmpty(t, order, "get order success")
}
