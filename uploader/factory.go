package uploader

import (
	"context"
	"github.com/lampnick/doctron/conf"
)

// new doctron
func NewDoctronUploader(ctx context.Context, doctronType string, uc UploadConfig) DoctronUploaderI {
	switch doctronType {
	case conf.DoctronUploaderAliOss:
		fac := &(aliossFactory{})
		return fac.createDoctronUploader(ctx, uc)
	case conf.DoctronUploaderMock:
		fac := &(mockFactory{})
		return fac.createDoctronUploader(ctx, uc)
	default:
		return nil
	}
}

type DoctronUploaderFactory interface {
	createDoctronUploader(uc UploadConfig) DoctronUploaderI
}

type aliossFactory struct {
}

func (ins *aliossFactory) createDoctronUploader(ctx context.Context, uc UploadConfig) DoctronUploaderI {
	return &AliOssUploader{
		DoctronUploader: DoctronUploader{
			ctx:          ctx,
			UploadConfig: uc,
		},
	}
}

type mockFactory struct {
}

func (ins *mockFactory) createDoctronUploader(ctx context.Context, uc UploadConfig) DoctronUploaderI {
	return &MockUploader{
		DoctronUploader: DoctronUploader{
			ctx:          ctx,
			UploadConfig: uc,
		},
	}
}
