package uploader

import (
	"context"
	"time"
)

type DoctronUploader struct {
	ctx context.Context
	UploadConfig
	uploadElapsed time.Duration
}

type DoctronUploaderI interface {
	Uploader
}
