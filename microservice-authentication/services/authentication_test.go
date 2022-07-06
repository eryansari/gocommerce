package services

import (
	"context"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gocommerce/microservice-authentication/pkg/model"
	"gocommerce/microservice-authentication/pkg/pbuff"
	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (r UserRepositoryMock) Create(context.Context, *model.UserAuth) (model.UserAuth, error) {
	args := r.Called()
	users := model.UserAuth{
		Username: "ex2",
		Email:    "ex2@ery.com",
		Phone:    "111",
	}
	return users, args.Error(1)
}

func (r UserRepositoryMock) CreateToken(context.Context, *model.UserToken) (model.UserToken, error) {
	args := r.Called()
	users := model.UserToken{
		Username: "ex",
	}
	return users, args.Error(1)
}

func (r UserRepositoryMock) DeleteTokenByUserAndDevice(context.Context, string, string) (bool, error) {
	args := r.Called()
	return args.Bool(0), args.Error(1)
}

func (r UserRepositoryMock) FindTokenByUserAndDevice(context.Context, string, string) (model.UserToken, error) {
	args := r.Called()
	ex := time.Now().Add(time.Hour * time.Duration(10))
	users := model.UserToken{
		Token:       "exToken",
		ExpiredDate: &ex,
	}
	return users, args.Error(1)
}

func (r UserRepositoryMock) FindByEmail(context.Context, string) (model.UserAuth, error) {
	args := r.Called()
	users := model.UserAuth{
		Email: "ex@ery.com",
	}
	return users, args.Error(1)
}

func (r UserRepositoryMock) FindByPhone(context.Context, string) (model.UserAuth, error) {
	args := r.Called()
	users := model.UserAuth{
		Phone: "111",
	}
	return users, args.Error(1)
}

func (r UserRepositoryMock) FindByUsername(context.Context, string) (model.UserAuth, error) {
	args := r.Called()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("111"), bcrypt.DefaultCost)
	users := model.UserAuth{
		Username: "ex",
		Password: string(hashedPassword),
		Status:   model.UserStatusActive,
	}
	return users, args.Error(1)
}

func (r UserRepositoryMock) FindByID(context.Context, string) (model.UserAuth, error) {
	args := r.Called()
	users := model.UserAuth{
		Username: "ex",
	}
	return users, args.Error(1)
}

func TestService_Login(t *testing.T) {
	repository := UserRepositoryMock{}
	repository.On("FindByUsername").Return(model.UserAuth{}, nil)
	repository.On("FindTokenByUserAndDevice").Return(model.UserAuth{}, nil)

	ctx := context.Background()

	service := AuthSvc{repository, logrus.StandardLogger()}
	req := &pbuff.LoginDto{
		Username: "ex",
		Password: "111",
		DeviceId: "111",
	}
	users, _ := service.Login(ctx, req)
	assert.Equal(t, users.Username, "ex")
}

func TestService_Register(t *testing.T) {
	repository := UserRepositoryMock{}
	repository.On("FindByUsername").Return(model.UserAuth{}, nil)
	repository.On("FindByPhone").Return(model.UserAuth{}, nil)
	repository.On("FindByEmail").Return(model.UserAuth{}, nil)
	repository.On("Create").Return(model.UserAuth{}, nil)

	ctx := context.Background()

	service := AuthSvc{repository, logrus.StandardLogger()}
	req := &pbuff.RegisterDto{
		Username: "ex2",
		Email:    "ex2@ery.com",
		Phone:    "222",
		Password: "222",
	}
	users, _ := service.Register(ctx, req)
	assert.Equal(t, users.Username, "ex2")
	assert.Equal(t, users.Phone, "111")
	assert.Equal(t, users.Email, "ex2@ery.com")
}
