package doctron_core

import (
	"github.com/lampnick/doctron/uploader"
	"time"
)

type pdf2Image struct {
	Doctron
}

func (p pdf2Image) Convert() ([]byte, error) {
	panic("implement me")
}

func (p pdf2Image) GetConvertElapsed() time.Duration {
	panic("implement me")
}

func (p pdf2Image) Upload(uf uploader.UploadFunc, uc uploader.UploadConfig) (string, error) {
	panic("implement me")
}

func (p pdf2Image) GetUploadElapsed() time.Duration {
	panic("implement me")
}
