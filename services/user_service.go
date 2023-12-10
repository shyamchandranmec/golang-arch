package services

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gofr.dev/pkg/gofr"

	"github.com/shyamchandranmec/golang-arch/db/dao"
)

type UserService struct {
	Q dao.UserQuerier
}

type UserServicer interface {
	CreateUser(c *gofr.Context, p CreateUserParams) (*dao.User, error)
}

type CreateUserParams struct {
	Name     string
	Username string
	Password string
}

func (us *UserService) CreateUser(c *gofr.Context, p CreateUserParams) (*dao.User, error) {
	u := dao.User{
		ID:        primitive.NewObjectID(),
		Name:      p.Name,
		Username:  p.Username,
		Password:  p.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Active:    true,
	}
	user, err := us.Q.CreateUser(c, &u)
	if err != nil {
		c.Logger.Error("unable to create user")
		u = dao.User{}
		return &u, err
	}
	c.Logger.Info("successfully added user ", user.Username)
	return &u, nil
}

func NewUserService() UserServicer {
	querier := &dao.UserQuery{}
	us := &UserService{
		Q: querier,
	}
	return us
}
