package routes

import (
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/muhammadikhsaan/simple-api-with-jwt.git/src/models"
	"github.com/muhammadikhsaan/simple-api-with-jwt.git/src/pages/auth/controller"
)

type context struct {
	i *iris.Application
}

func New() models.RouterModel {
	return &context{
		i: iris.Default(),
	}
}

func (ctx *context) Run() {
	ctx.initialize()
	ctx.routes()
	ctx.i.Listen(":1000")
}

func (ctx *context) initialize() {
	f, _ := os.Create("log/access.log")
	ctx.i.Logger().SetOutput(f)
	ctx.i.Use(iris.Compression)
	ctx.i.Validator = validator.New()
}

func (ctx *context) routes() {
	auth := ctx.i.Party("/auth")
	{
		auth.Post("/login", controller.Login)
		auth.Post("/register", controller.Register)
		auth.Get("/validate", controller.Validate)
	}
}
