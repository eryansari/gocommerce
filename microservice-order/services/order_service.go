package services

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/sirupsen/logrus"
	"gocommerce/microservice-order/pkg/model"
	"gocommerce/microservice-order/pkg/pbuff"
	"gocommerce/microservice-order/repositories"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// OrderService
type OrderSvc struct {
	orepo repositories.OrderRepository
	log   *logrus.Logger
}

// NewOrderSvc
func NewOrderSvc(orepo repositories.OrderRepository, log *logrus.Logger) *OrderSvc {
	return &OrderSvc{
		orepo: orepo,
		log:   log,
	}
}

// CreateOrder
func (s *OrderSvc) Create(ctx context.Context, ord *pbuff.OrderReqDto) (*pbuff.OrderDto, error) {
	var (
		now            = time.Now()
		ordDto         *pbuff.OrderDto
		ordDetailSlice []*pbuff.OrderDetailDto
		ordModelSlice  []model.OrderDetail
	)

	if len(ord.Order) == 0 {
		return nil, status.Errorf(codes.NotFound, "order required")
	}

	if len(ord.Order) > 1 {
		return nil, status.Errorf(codes.NotFound, "order must be only 1 data")
	}

	for _, p := range ord.Order {
		if len(p.ProductId) == 0 {
			return nil, status.Errorf(codes.NotFound, "product_id required")
		}
		if p.Qty == 0 {
			return nil, status.Errorf(codes.NotFound, "qty required")
		}
	}

	rand.Seed(time.Now().UnixNano())
	noInv := fmt.Sprintf("INV/%v/%v", now.Format("060201150405"), rand.Intn(9999-1000)+1000)

	om := &model.Order{
		NoInvoice: noInv,
		Status:    false,
	}

	for _, o := range ord.Order {
		od := model.OrderDetail{
			ProductID: o.ProductId,
			Qty:       o.Qty,
		}
		ordModelSlice = append(ordModelSlice, od)
	}
	om.OrderDetail = ordModelSlice

	ores, err := s.orepo.Create(ctx, om)
	if err != nil {
		s.log.Error(err)
		return nil, status.Errorf(codes.Internal, "Create order failed")
	}

	pid, _ := hex.DecodeString(ores.ID.Hex())
	ordDto = &pbuff.OrderDto{
		Id:          string(pid),
		UserId:      ores.UserID,
		NoInv:       ores.NoInvoice,
		Status:      ores.Status,
		OrderDetail: ordDetailSlice,
	}

	for _, o := range ores.OrderDetail {
		od := pbuff.OrderDetailDto{
			ProductId: o.ProductID,
			Qty:       o.Qty,
		}
		ordDetailSlice = append(ordDetailSlice, &od)
	}
	ordDto.OrderDetail = ordDetailSlice

	s.log.Debugf("created Order(ID: %s, No Invoice: %s)", ordDto.Id, ordDto.NoInv)

	return ordDto, nil
}

// GetOrderByID
func (s *OrderSvc) GetOrderByID(ctx context.Context, id *wrappers.StringValue) (*pbuff.OrderDto, error) {
	pdetl, err := s.orepo.FindByID(id.GetValue())
	var productID string
	var qty int64
	if err != nil {
		s.log.Error(err)
		return nil, status.Errorf(codes.NotFound, "Order not found")
	}
	for _, p := range pdetl.OrderDetail {
		productID = p.ProductID
		qty = p.Qty
	}
	pid, _ := hex.DecodeString(pdetl.ID.Hex())
	response := &pbuff.OrderDto{
		Id:     string(pid),
		UserId: pdetl.UserID,
		NoInv:  pdetl.NoInvoice,
		Status: pdetl.Status,
		OrderDetail: []*pbuff.OrderDetailDto{
			{
				ProductId: productID,
				Qty:       qty,
			},
		},
	}

	return response, nil

}

// ListOrderByUserID
func (s *OrderSvc) ListOrderByUserID(ctx context.Context, userId *wrappers.StringValue) (*pbuff.OrderDtoList, error) {
	var (
		plist   *pbuff.OrderDtoList
		ordDtos []*pbuff.OrderDto
		ordDto  *pbuff.OrderDto
	)

	orders, err := s.orepo.FindByUserID(userId.GetValue())
	if err != nil {
		s.log.Error(err)
		return nil, status.Errorf(codes.NotFound, "Order not found")
	}

	for _, p := range orders {
		var productID string
		var qty int64
		pid, _ := hex.DecodeString(p.ID.Hex())
		for _, q := range p.OrderDetail {
			productID = q.ProductID
			qty = q.Qty
		}
		ordDto = &pbuff.OrderDto{
			Id:     string(pid),
			UserId: userId.GetValue(),
			NoInv:  p.NoInvoice,
			Status: p.Status,
			OrderDetail: []*pbuff.OrderDetailDto{
				{
					ProductId: productID,
					Qty:       qty,
				},
			},
		}
		ordDtos = append(ordDtos, ordDto)
	}

	plist = new(pbuff.OrderDtoList)
	plist.List = ordDtos

	if len(plist.List) == 0 {
		return nil, status.Errorf(codes.NotFound, "Order not found")
	}

	return plist, nil
}
