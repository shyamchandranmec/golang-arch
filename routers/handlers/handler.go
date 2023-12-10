package handlers

import (
	"fmt"

	"github.com/artyom/httpflags"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/template"

	"github.com/shyamchandranmec/golang-arch/routers/models"
	"github.com/shyamchandranmec/golang-arch/services"
)

func Index(c *gofr.Context) (interface{}, error) {
	return template.Template{
		Directory: "html",
		File:      "index.html",
		Type:      template.HTML,
	}, nil
}

func Register(c *gofr.Context) (interface{}, error) {
	c.Logger.Info("Request to register user ")
	var reg models.RegisterReq
	err := c.Bind(&reg)
	if err != nil {
		c.Logger.Errorf("Error in binding request body %w ", err)
	}
	c.Logger.Info("Registering user ", reg.Name)

	params := services.CreateUserParams{
		Name:     reg.Name,
		Password: reg.Password,
		Username: reg.Username,
	}
	us := services.NewUserService()
	user, err := us.CreateUser(c, params)
	// status := c.Redis.Set(c, "user", reg.Name, 1*time.Minute)
	// if status.Err() != nil {
	// 	c.Logger.Errorf("error  is %w", status.Err())
	// }
	if err != nil {
		return "Failed", err
	}
	return user, nil
}

func Login(c *gofr.Context) (interface{}, error) {
	c.Logger.Info("Login request ")
	var login models.LoginReq
	x := &models.LoginReq{}
	err := c.Bind(&login)
	err2 := httpflags.Parse(x, c.Request())
	if err2 != nil {
		c.Logger.Errorf("Error in parse form body %w", &err2)
	}
	fmt.Println("value of login is ", login)
	if err != nil {
		c.Logger.Errorf("Error in binding request body %w ", err)
		fmt.Println("Reached error")

	}
	c.Logger.Info("Login request for username ", login.Username)
	return login.Username, nil
}
