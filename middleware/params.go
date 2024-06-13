package middleware

import (
	"net/url"

	"github.com/kataras/iris/v12"
	"github.com/lampnick/doctron/common"
)

func CheckParams(ctx iris.Context) {
	webUrl := ctx.URLParam("url")
	if webUrl == "" {
		outputDTO := common.NewDefaultOutputDTO(nil)
		outputDTO.Code = common.InvalidUrl
		_, _ = common.NewJsonOutput(ctx, outputDTO)
		return
	}

	u, err := url.Parse(webUrl)
	if err != nil {
		outputDTO := common.NewDefaultOutputDTO(nil)
		outputDTO.Code = common.InvalidUrl
		outputDTO.Message = err.Error()
		_, _ = common.NewJsonOutput(ctx, outputDTO)
		return
	}

	if u.Scheme != "http" && u.Scheme != "https" && u.Scheme != "file" {
		outputDTO := common.NewDefaultOutputDTO(nil)
		outputDTO.Code = common.InvalidUrlScheme
		_, _ = common.NewJsonOutput(ctx, outputDTO)
		return
	}

	ctx.Next()
}
