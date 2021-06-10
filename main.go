package main

import (
	"github.com/muhammadikhsaan/simple-api-with-jwt.git/src/routes"
)

var (
	app = routes.New()
)

func main() {
	defer app.Run()
}
