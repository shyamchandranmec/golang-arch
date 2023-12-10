package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gofr.dev/pkg/gofr"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	Active    bool               `json:"active" bson:"active"`
}

type UserQuery struct {
}

type UserQuerier interface {
	CreateUser(c *gofr.Context, u *User) (*User, error)
}

func (q *UserQuery) CreateUser(c *gofr.Context, u *User) (*User, error) {
	col := c.MongoDB.Collection("users")
	res, err := col.InsertOne(c, u)
	if err != nil {
		c.Logger.Errorf("unable to create user %w", err)
	}
	c.Logger.Info("user created ", res)
	return u, err
}
