package main

import (
	"gofr.dev/pkg/gofr"

	"github.com/shyamchandranmec/golang-arch/routers"
)

func main() {
	app := gofr.New()
	routers.AddRoutes(app)
	app.Start()
}
