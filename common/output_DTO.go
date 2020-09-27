package common

import (
	"github.com/kataras/iris/v12"
)

type OutputDTO struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewDefaultOutputDTO(data interface{}) OutputDTO {
	return OutputDTO{
		Code:    Success,
		Message: ErrMsg[Success],
		Data:    data,
	}
}

func NewJsonOutput(ctx iris.Context, outputDTO OutputDTO) (int, error) {
	if outputDTO.Code != Success && outputDTO.Message == "" {
		msg, ok := ErrMsg[outputDTO.Code]
		if ok {
			outputDTO.Message = msg
		}
	}
	return ctx.JSON(outputDTO)
}
