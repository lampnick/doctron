package converter

import (
	"context"
	uuid "github.com/iris-contrib/go.uuid"
	irisContext "github.com/kataras/iris/v12/context"
)

type ConvertConfig struct {
	//the source document's url to be converted.
	Url       string `validate:"required,url"`
	UploadKey string `validate:"omitempty"`
	Params    Params
}

type DoctronConfig struct {
	TraceId     uuid.UUID
	IrisCtx     irisContext.Context
	Ctx         context.Context
	DoctronType int
	ConvertConfig
}

// convert params
type Params interface {
}
