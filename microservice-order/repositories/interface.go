package repositories

import (
	"context"

	"gocommerce/microservice-order/pkg/model"
)

type (
	Repository interface {
		OrderRepository() OrderRepository
	}

	// OrderRepo
	OrderRepository interface {
		Create(c context.Context, t *model.Order) (model.Order, error)
		FindByID(id string) (model.Order, error)
		FindByUserID(userID string) ([]model.Order, error)
	}
)
