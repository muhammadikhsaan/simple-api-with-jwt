package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/muhammadikhsaan/simple-api-with-jwt.git/src/helpers/validation"
	"github.com/muhammadikhsaan/simple-api-with-jwt.git/src/pages/auth/models"
	"github.com/muhammadikhsaan/simple-api-with-jwt.git/src/pages/auth/usecase"
)

func Register(ctx iris.Context) {
	var register models.Users

	if err := ctx.ReadBody(&register); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			// Wrap the errors with JSON format, the underline library returns the errors as interface.
			validationErrors := validation.ValidationErrors(errs)

			// Fire an application/json+problem response and stop the handlers chain.
			ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
				Title("Validation error").
				Detail("One or more fields failed to be validated").
				Type("/auth/register").
				Key("errors", validationErrors))

			return
		}

		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	result, err := usecase.IsRegister(&register)

	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	ctx.StopWithJSON(200, result)
}
