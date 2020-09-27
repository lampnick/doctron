package uploader

import (
	"github.com/lampnick/doctron/conf"
	"time"
)

type MockUploader struct {
	DoctronUploader
}

func (ins *MockUploader) Upload() (url string, err error) {
	start := time.Now()
	defer func() {
		ins.uploadElapsed = time.Since(start)
	}()

	return "http://" + conf.LoadedConfig.Oss.PrivateServerDomain + "/" + ins.UploadConfig.Key, nil
}

func (ins *MockUploader) GetUploadElapsed() time.Duration {
	return ins.uploadElapsed
}
