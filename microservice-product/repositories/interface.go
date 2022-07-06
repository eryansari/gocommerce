package repositories

import (
	"context"

	"gocommerce/microservice-product/pkg/model"
)

type (
	Repository interface {
		ProductRepository() ProductRepository
	}

	// Product Repo
	ProductRepository interface {
		Create(c context.Context, t *model.Product) (model.Product, error)
		FindByID(id string) (model.Product, error)
		FindAll() ([]model.Product, error)
	}
)
