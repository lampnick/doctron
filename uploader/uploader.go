package uploader

import (
	"time"
)

type UploadFunc func(uc UploadConfig) (url string, err error)

//go:generate mockgen -source=./uploader.go -destination=./mock/uploader_mock.go -package=mock_uploader

type Uploader interface {
	Upload() (string, error)
	GetUploadElapsed() time.Duration
}
