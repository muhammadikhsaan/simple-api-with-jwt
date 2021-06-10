package controller

import (
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/muhammadikhsaan/simple-api-with-jwt.git/src/helpers/tokens"
	"github.com/muhammadikhsaan/simple-api-with-jwt.git/src/helpers/validation"
	"github.com/muhammadikhsaan/simple-api-with-jwt.git/src/pages/auth/models"
	"github.com/muhammadikhsaan/simple-api-with-jwt.git/src/pages/auth/usecase"
)

func Login(ctx iris.Context) {
	var login models.Users

	if err := ctx.ReadBody(&login); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			// Wrap the errors with JSON format, the underline library returns the errors as interface.
			validationErrors := validation.ValidationErrors(errs)

			// Fire an application/json+problem response and stop the handlers chain.
			ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
				Title("Validation error").
				Detail("One or more fields failed to be validated").
				Type("/auth/login").
				Key("errors", validationErrors))

			return
		}

		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	is, err := usecase.IsLogin(&login)

	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	if !is {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	token, err := jwtToken.GenerateToken(tokens.SignetureToken{
		Email: login.Email,
	})

	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	ctx.SetCookie(&http.Cookie{
		Name:     "X-JWT-TOKEN",
		Value:    token,
		Path:     "/",
		Domain:   "localhost",
		Expires:  time.Now().Add(365 * 24 * time.Hour),
		Secure:   true,
		HttpOnly: true,
	})

	ctx.StopWithJSON(200, models.LoginResponse{
		Token: token,
	})
}
