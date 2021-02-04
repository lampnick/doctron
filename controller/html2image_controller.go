package controller

import (
	"context"
	"errors"
	"time"

	"github.com/gorilla/schema"
	uuid "github.com/iris-contrib/go.uuid"
	"github.com/kataras/iris/v12"
	irisContext "github.com/kataras/iris/v12/context"
	"github.com/lampnick/doctron/common"
	"github.com/lampnick/doctron/conf"
	"github.com/lampnick/doctron/converter"
	"github.com/lampnick/doctron/converter/doctron_core"
	"github.com/lampnick/doctron/worker"
	"gopkg.in/go-playground/validator.v9"
)

func Html2ImageHandler(ctx iris.Context) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Duration(conf.LoadedConfig.Doctron.ConvertTimeout)*time.Second)
	defer cancel()
	traceId, _ := uuid.NewV4()
	outputDTO := common.NewDefaultOutputDTO(nil)
	doctronConfig, err := initHtml2ImageConfig(ctx)
	if err != nil {
		outputDTO.Code = common.InvalidParams
		outputDTO.Message = err.Error()
		_, _ = common.NewJsonOutput(ctx, outputDTO)
		return
	}
	doctronConfig.TraceId = traceId
	doctronConfig.Ctx = ctxTimeout
	doctronConfig.IrisCtx = ctx
	doctronConfig.DoctronType = doctron_core.DoctronHtml2Image

	doctronOutputDTO, err := worker.Pool.ProcessTimed(doctronConfig, time.Duration(conf.LoadedConfig.Doctron.ConvertTimeout)*time.Second)
	if err != nil {
		outputDTO.Code = common.ConvertHtml2ImageFailed
		outputDTO.Message = "worker run process failed." + err.Error()
		_, _ = common.NewJsonOutput(ctx, outputDTO)
		return
	}
	doctronOutput, ok := doctronOutputDTO.(worker.DoctronOutputDTO)
	if !ok {
		outputDTO.Code = common.ConvertHtml2ImageFailed
		outputDTO.Message = "error type assert to DoctronOutputDTO"
		_, _ = common.NewJsonOutput(ctx, outputDTO)
		return
	}

	if errors.Is(doctronOutput.Err, worker.ErrNoNeedToUpload) {
		ctx.Header(irisContext.ContentTypeHeaderKey, "image/png")
		_, err = ctx.Write(doctronOutput.Buf)
		if err != nil {
			outputDTO.Code = common.ConvertHtml2ImageFailed
			_, _ = common.NewJsonOutput(ctx, outputDTO)
			return
		}
	}
	if doctronOutput.Err != nil {
		outputDTO.Code = common.ConvertHtml2ImageFailed
		outputDTO.Message = doctronOutput.Err.Error()
		_, _ = common.NewJsonOutput(ctx, outputDTO)
		return
	}
	outputDTO.Data = doctronOutput.Url
	_, _ = common.NewJsonOutput(ctx, outputDTO)
	return
}

func initHtml2ImageConfig(ctx iris.Context) (converter.DoctronConfig, error) {
	result := converter.DoctronConfig{}
	requestDTO := newDefaultHtml2ImageRequestDTO()
	// decode query params
	var decoder = schema.NewDecoder()
	err := decoder.Decode(requestDTO, ctx.Request().URL.Query())
	if err != nil {
		log(ctx, "[%s][%s]", ctx.Request().RequestURI, err)
	}
	// validate
	validate := validator.New()
	err = validate.Struct(requestDTO)
	if err != nil {
		return result, err
	}
	result.Url = requestDTO.Url
	result.UploadKey = requestDTO.UploadKey
	result.Params = convertToHtml2ImageParams(requestDTO)
	return result, nil
}

func convertToHtml2ImageParams(requestDTO *Html2ImageRequestDTO) doctron_core.Html2ImageParams {
	params := doctron_core.NewDefaultHtml2ImageParams()
	params.Format = requestDTO.Format
	params.Quality = requestDTO.Quality
	params.CustomClip = requestDTO.CustomClip
	params.Clip.X = requestDTO.ClipX
	params.Clip.Y = requestDTO.ClipY
	params.Clip.Width = requestDTO.ClipWidth
	params.Clip.Height = requestDTO.ClipHeight
	params.Clip.Scale = requestDTO.ClipScale
	params.FromSurface = requestDTO.FromSurface
	params.WaitingTime = requestDTO.WaitingTime
	return params
}

func newDefaultHtml2ImageRequestDTO() *Html2ImageRequestDTO {
	return &Html2ImageRequestDTO{
		Format:      doctron_core.FormatPng,
		Quality:     doctron_core.DefaultQuality,
		CustomClip:  false,
		ClipX:       doctron_core.DefaultViewportX,
		ClipY:       doctron_core.DefaultViewportY,
		ClipWidth:   doctron_core.DefaultViewportWidth,
		ClipHeight:  doctron_core.DefaultViewportHeight,
		ClipScale:   doctron_core.DefaultViewportScale,
		FromSurface: doctron_core.DefaultFromSurface,
		WaitingTime: doctron_core.DefaultWaitingTime,
	}
}
