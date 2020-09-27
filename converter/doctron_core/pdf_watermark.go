package doctron_core

import (
	"github.com/lampnick/doctron/uploader"
	"time"
)

type pdfWatermark struct {
	Doctron
}

func (p pdfWatermark) Convert() ([]byte, error) {
	panic("implement me")
}

func (p pdfWatermark) GetConvertElapsed() time.Duration {
	panic("implement me")
}

func (p pdfWatermark) Upload(uf uploader.UploadFunc, uc uploader.UploadConfig) (string, error) {
	panic("implement me")
}

func (p pdfWatermark) GetUploadElapsed() time.Duration {
	panic("implement me")
}
