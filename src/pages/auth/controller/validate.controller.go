package controller

import (
	"github.com/kataras/iris/v12"
)

func Validate(ctx iris.Context) {
	token := ctx.GetCookie("X-JWT-TOKEN")

	claims, err := jwtToken.ValidateToken(token)

	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	ctx.StopWithJSON(iris.StatusOK, claims)
}
