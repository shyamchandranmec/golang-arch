package routers

import (
	"github.com/shyamchandranmec/golang-arch/routers/handlers"
	"gofr.dev/pkg/gofr"
)

func AddRoutes(app *gofr.Gofr) {
	app.GET("/", handlers.Index)
	app.POST("/register", handlers.Register)
	app.POST("/login", handlers.Login)
}
