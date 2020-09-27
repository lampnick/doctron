package alioss

import (
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

//OSS config
type OssConfig struct {
	Endpoint                 string
	AccessKeyId              string
	AccessKeySecret          string
	BucketName               string
	PrivateImageServerDomain string
}
type OssHelper struct {
	client *oss.Client
	config OssConfig
}

func NewOssHelper(c OssConfig, options ...oss.ClientOption) (*OssHelper, error) {
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
	return fmt.Sprintf("https://%s/%s", h.config.PrivateImageServerDomain, objectKey), nil
}
