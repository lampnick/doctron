package uploader

import (
	"errors"
	"github.com/lampnick/doctron/conf"
	"github.com/lampnick/doctron/pkg/alioss"
	"time"
)

type AliOssUploader struct {
	DoctronUploader
}

var ErrNoNeedToUpload = errors.New("no need to upload")

func (ins *AliOssUploader) Upload() (url string, err error) {
	if ins.Key == "" {
		return "", ErrNoNeedToUpload
	}
	start := time.Now()
	defer func() {
		ins.uploadElapsed = time.Since(start)
	}()
	helper, err := alioss.NewOssHelper(conf.OssConfig)
	if err != nil {
		return "", err
	}
	uploadUrl, err := helper.Upload(ins.Key, ins.Stream)
	if err != nil {
		return "", err
	}
	return uploadUrl, nil
}

func (ins *AliOssUploader) GetUploadElapsed() time.Duration {
	return ins.uploadElapsed
}
