package controller

import (
	"github.com/kataras/iris/v12"
)

func Pdf2ImageHandler(ctx iris.Context) {
	ctx.Application().Logger().Infof("controller: %s", "Pdf2ImageHandler")

}
