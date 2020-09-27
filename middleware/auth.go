package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/lampnick/doctron/common"
	"github.com/lampnick/doctron/conf"
)

func AuthMiddleware(ctx iris.Context) {
	username := ctx.URLParam("u")
	if username == "" {
		username = ctx.URLParam("username")

	}
	password := ctx.URLParam("p")
	if password == "" {
		password = ctx.URLParam("password")

	}
	valid := checkAuth(username, password)
	if valid == false {
		outputDTO := common.NewDefaultOutputDTO(nil)
		outputDTO.Code = common.AuthFailed
		_, _ = common.NewJsonOutput(ctx, outputDTO)
		return
	}

	ctx.Next()
}

func checkAuth(username, password string) bool {
	for _, user := range conf.LoadedConfig.Doctron.User {
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}
