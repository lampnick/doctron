package alioss

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"gopkg.in/go-playground/validator.v9"
)

//OSS config
type OssConfig struct {
	Endpoint            string `validate:"required"`
	AccessKeyId         string `validate:"required"`
	AccessKeySecret     string `validate:"required"`
	BucketName          string `validate:"required"`
	PrivateServerDomain string `validate:"required"`
}
type OssHelper struct {
	client *oss.Client
	config OssConfig
}

func NewOssHelper(c OssConfig, options ...oss.ClientOption) (*OssHelper, error) {
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		return nil, errors.New("alioss uploader config not set")
	}
	client, e := oss.New(c.Endpoint, c.AccessKeyId, c.AccessKeySecret, options...)
	if e != nil {
		return nil, e
	}
	return &OssHelper{
		client: client,
		config: c,
	}, nil
}

func (h OssHelper) Upload(objectKey string, b []byte, options ...oss.Option) (url string, err error) {
	// 获取存储空间。
	bucket, err := h.client.Bucket(h.config.BucketName)
	if err != nil {
		return "", err
	}
	// 上传文件。
	reader := bytes.NewReader(b)
	err = bucket.PutObject(objectKey, reader, options...)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://%s/%s", h.config.PrivateServerDomain, objectKey), nil
}
