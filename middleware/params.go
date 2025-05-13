package middleware

import (
	"net/url"

	"github.com/gorilla/schema"
	"github.com/kataras/iris/v12"
	"github.com/lampnick/doctron/common"
)

func CheckParams(ctx iris.Context) {
	type Param struct {
		Url string `schema:"url,omitempty" validate:"required,url"`
	}
	dto := Param{}
	decoder := schema.NewDecoder()
	err := decoder.Decode(&dto, ctx.Request().URL.Query()) 
	if err != nil {
		outputDTO := common.NewDefaultOutputDTO(nil)
		outputDTO.Code = common.InvalidUrl
		outputDTO.Message = err.Error()
		_, _ = common.NewJsonOutput(ctx, outputDTO)
		return
	}
	webUrl := dto.Url
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

	if u.Scheme != "http" && u.Scheme != "https" {
		outputDTO := common.NewDefaultOutputDTO(nil)
		outputDTO.Code = common.InvalidUrlScheme
		_, _ = common.NewJsonOutput(ctx, outputDTO)
		return
	}

	ctx.Next()
}
