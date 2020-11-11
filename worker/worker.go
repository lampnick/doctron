package worker

import (
	"errors"

	"github.com/Jeffail/tunny"
	"github.com/kataras/iris/v12"
	"github.com/lampnick/doctron/conf"
	"github.com/lampnick/doctron/converter"
	"github.com/lampnick/doctron/converter/doctron_core"
	"github.com/lampnick/doctron/uploader"
)

var Pool *tunny.Pool

var ErrNoNeedToUpload = errors.New("no need to upload")
var (
	ErrWrongDoctronParam = errors.New("wrong doctron params given")
)

func log(ctx iris.Context, format string, args ...interface{}) {
	ctx.Application().Logger().Infof(format, args...)
}

type DoctronOutputDTO struct {
	Buf []byte
	Url string
	Err error
}

func DoctronHandler(params interface{}) interface{} {
	doctronOutputDTO := DoctronOutputDTO{}
	doctronConfig, ok := params.(converter.DoctronConfig)
	if !ok {
		doctronOutputDTO.Err = ErrWrongDoctronParam
		return doctronOutputDTO
	}

	doctron := doctron_core.NewDoctron(doctronConfig.Ctx, doctronConfig.DoctronType, doctronConfig.ConvertConfig)

	convertBytes, err := doctron.Convert()
	log(doctronConfig.IrisCtx, "uuid:[%s],doctron.Convert Elapsed [%s],url:[%s]", doctronConfig.TraceId, doctron.GetConvertElapsed(), doctronConfig.IrisCtx.Request().RequestURI)
	if err != nil {
		doctronOutputDTO.Err = err
		return doctronOutputDTO
	}

	doctronOutputDTO.Buf = convertBytes
	if doctronConfig.UploadKey == "" {
		doctronOutputDTO.Err = ErrNoNeedToUpload
		return doctronOutputDTO
	} else {
		doctronUploader := uploader.NewDoctronUploader(
			doctronConfig.Ctx,
			conf.LoadedConfig.Doctron.Uploader,
			uploader.UploadConfig{Key: doctronConfig.UploadKey, Stream: convertBytes},
		)
		uploadUrl, err := doctronUploader.Upload()
		log(doctronConfig.IrisCtx, "uuid:[%s],doctron.Upload Elapsed [%s],url:[%s]", doctronConfig.TraceId, doctronUploader.GetUploadElapsed(), doctronConfig.IrisCtx.Request().RequestURI)
		if err != nil {
			doctronOutputDTO.Err = err
			return doctronOutputDTO
		}
		doctronOutputDTO.Url = uploadUrl
		return doctronOutputDTO
	}
}
