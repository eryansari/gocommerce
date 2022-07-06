package repositories

import (
	"context"

	"gocommerce/microservice-authentication/pkg/model"
)

type (
	Repository interface {
		UserRepository() UserRepository
	}

	// User Interface Repo
	UserRepository interface {
		Create(c context.Context, t *model.UserAuth) (model.UserAuth, error)
		CreateToken(c context.Context, t *model.UserToken) (model.UserToken, error)
		DeleteTokenByUserAndDevice(c context.Context, username, deviceID string) (bool, error)
		FindTokenByUserAndDevice(c context.Context, username, deviceID string) (model.UserToken, error)
		FindByEmail(c context.Context, email string) (model.UserAuth, error)
		FindByPhone(c context.Context, phone string) (model.UserAuth, error)
		FindByUsername(c context.Context, uname string) (model.UserAuth, error)
		FindByID(c context.Context, id string) (model.UserAuth, error)
	}
)
