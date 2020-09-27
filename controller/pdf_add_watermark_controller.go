package controller

import (
	"github.com/kataras/iris/v12"
)

func PdfAddWatermarkHandler(ctx iris.Context) {
	ctx.Application().Logger().Infof("controller: %s", "PdfAddWatermarkHandler")

}
