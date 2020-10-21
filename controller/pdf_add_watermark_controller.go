package controller

import (
	"context"
	"errors"
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
	"time"
)

func PdfAddWatermarkHandler(ctx iris.Context) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Duration(conf.LoadedConfig.Doctron.ConvertTimeout)*time.Second)
	defer cancel()
	traceId, _ := uuid.NewV4()
	outputDTO := common.NewDefaultOutputDTO(nil)
	doctronConfig, err := initPdfWatermarkConfig(ctx)
	if err != nil {
		outputDTO.Code = common.InvalidParams
		outputDTO.Message = err.Error()
		_, _ = common.NewJsonOutput(ctx, outputDTO)
		return
	}
	doctronConfig.TraceId = traceId
	doctronConfig.Ctx = ctxTimeout
	doctronConfig.IrisCtx = ctx
	doctronConfig.DoctronType = doctron_core.DoctronPdfWatermark

	doctronOutputDTO, err := worker.Pool.ProcessTimed(doctronConfig, time.Duration(conf.LoadedConfig.Doctron.ConvertTimeout)*time.Second)
	if err != nil {
		outputDTO.Code = common.ConvertPdfWatermarkFailed
		outputDTO.Message = "worker run process failed." + err.Error()
		_, _ = common.NewJsonOutput(ctx, outputDTO)
		return
	}
	doctronOutput, ok := doctronOutputDTO.(worker.DoctronOutputDTO)
	if !ok {
		outputDTO.Code = common.ConvertPdfWatermarkFailed
		outputDTO.Message = "error type assert to DoctronOutputDTO"
		_, _ = common.NewJsonOutput(ctx, outputDTO)
		return
	}
	if errors.Is(doctronOutput.Err, worker.ErrNoNeedToUpload) {
		ctx.Header(irisContext.ContentTypeHeaderKey, "application/pdf")
		_, err = ctx.Write(doctronOutput.Buf)
		if err != nil {
			outputDTO.Code = common.ConvertPdfWatermarkFailed
			_, _ = common.NewJsonOutput(ctx, outputDTO)
			return
		}
	}
	if doctronOutput.Err != nil {
		outputDTO.Code = common.ConvertPdfWatermarkFailed
		outputDTO.Message = doctronOutput.Err.Error()
		_, _ = common.NewJsonOutput(ctx, outputDTO)
		return
	}
	outputDTO.Data = doctronOutput.Url
	_, _ = common.NewJsonOutput(ctx, outputDTO)
	return
}

func initPdfWatermarkConfig(ctx iris.Context) (converter.DoctronConfig, error) {
	result := converter.DoctronConfig{}
	requestDTO := newDefaultPdfWatermarkRequestDTO()
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
	result.Params = pdfWatermarkParams(requestDTO)
	return result, nil
}

func pdfWatermarkParams(requestDTO *PdfWatermarkRequestDTO) doctron_core.PdfWatermarkParams {
	params := doctron_core.NewDefaultPdfWatermarkParams()
	params.ImageUrl = requestDTO.ImageUrl
	params.WatermarkType = requestDTO.WatermarkType
	return params
}

func newDefaultPdfWatermarkRequestDTO() *PdfWatermarkRequestDTO {
	return &PdfWatermarkRequestDTO{
		WatermarkType: 0,
		ImageUrl:      "",
	}
}
