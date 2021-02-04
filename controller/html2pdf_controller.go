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

func Html2PdfHandler(ctx iris.Context) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Duration(conf.LoadedConfig.Doctron.ConvertTimeout)*time.Second)
	defer cancel()
	traceId, _ := uuid.NewV4()
	outputDTO := common.NewDefaultOutputDTO(nil)
	doctronConfig, err := initHtml2PdfConfig(ctx)
	if err != nil {
		outputDTO.Code = common.InvalidParams
		outputDTO.Message = err.Error()
		_, _ = common.NewJsonOutput(ctx, outputDTO)
		return
	}
	doctronConfig.TraceId = traceId
	doctronConfig.Ctx = ctxTimeout
	doctronConfig.IrisCtx = ctx
	doctronConfig.DoctronType = doctron_core.DoctronHtml2Pdf

	doctronOutputDTO, err := worker.Pool.ProcessTimed(doctronConfig, time.Duration(conf.LoadedConfig.Doctron.ConvertTimeout)*time.Second)
	if err != nil {
		outputDTO.Code = common.ConvertPdfFailed
		outputDTO.Message = "worker run process failed." + err.Error()
		_, _ = common.NewJsonOutput(ctx, outputDTO)
		return
	}
	doctronOutput, ok := doctronOutputDTO.(worker.DoctronOutputDTO)
	if !ok {
		outputDTO.Code = common.ConvertPdfFailed
		outputDTO.Message = "error type assert to DoctronOutputDTO"
		_, _ = common.NewJsonOutput(ctx, outputDTO)
		return
	}

	if errors.Is(doctronOutput.Err, worker.ErrNoNeedToUpload) {
		ctx.Header(irisContext.ContentTypeHeaderKey, "application/pdf")
		_, err = ctx.Write(doctronOutput.Buf)
		if err != nil {
			outputDTO.Code = common.ConvertPdfFailed
			_, _ = common.NewJsonOutput(ctx, outputDTO)
			return
		}
	}
	if doctronOutput.Err != nil {
		outputDTO.Code = common.ConvertPdfFailed
		outputDTO.Message = doctronOutput.Err.Error()
		_, _ = common.NewJsonOutput(ctx, outputDTO)
		return
	}
	outputDTO.Data = doctronOutput.Url
	_, _ = common.NewJsonOutput(ctx, outputDTO)
	return
}

func initHtml2PdfConfig(ctx iris.Context) (converter.DoctronConfig, error) {
	result := converter.DoctronConfig{}
	requestDTO := newDefaultHtml2PdfRequestDTO()
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
	result.Params = convertToPDFParams(requestDTO)
	return result, nil
}

func convertToPDFParams(requestDTO *Html2PdfRequestDTO) doctron_core.PDFParams {
	params := doctron_core.NewDefaultPDFParams()
	params.Landscape = requestDTO.Landscape
	params.DisplayHeaderFooter = requestDTO.DisplayHeaderFooter
	params.PrintBackground = requestDTO.PrintBackground
	params.Scale = requestDTO.Scale
	params.PaperWidth = requestDTO.PaperWidth
	params.PaperHeight = requestDTO.PaperHeight
	params.MarginTop = requestDTO.MarginTop
	params.MarginBottom = requestDTO.MarginBottom
	params.MarginLeft = requestDTO.MarginLeft
	params.MarginRight = requestDTO.MarginRight
	params.PageRanges = requestDTO.PageRanges
	params.IgnoreInvalidPageRanges = requestDTO.IgnoreInvalidPageRanges
	params.HeaderTemplate = requestDTO.HeaderTemplate
	params.FooterTemplate = requestDTO.FooterTemplate
	params.PreferCSSPageSize = requestDTO.PreferCSSPageSize
	params.WaitingTime = requestDTO.WaitingTime
	return params
}

func newDefaultHtml2PdfRequestDTO() *Html2PdfRequestDTO {
	return &Html2PdfRequestDTO{
		Landscape:               doctron_core.DefaultLandscape,
		DisplayHeaderFooter:     doctron_core.DefaultDisplayHeaderFooter,
		PrintBackground:         doctron_core.DefaultPrintBackground,
		Scale:                   doctron_core.DefaultScale,
		PaperWidth:              doctron_core.DefaultPaperWidth,
		PaperHeight:             doctron_core.DefaultPaperHeight,
		MarginTop:               doctron_core.DefaultMarginTop,
		MarginBottom:            doctron_core.DefaultMarginBottom,
		MarginLeft:              doctron_core.DefaultMarginLeft,
		MarginRight:             doctron_core.DefaultMarginRight,
		PageRanges:              doctron_core.DefaultPageRanges,
		IgnoreInvalidPageRanges: doctron_core.DefaultIgnoreInvalidPageRanges,
		HeaderTemplate:          doctron_core.DefaultHeaderTemplate,
		FooterTemplate:          doctron_core.DefaultFooterTemplate,
		PreferCSSPageSize:       doctron_core.DefaultPreferCSSPageSize,
		WaitingTime:             doctron_core.DefaultWaitingTime,
	}
}
