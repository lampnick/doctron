package conf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lampnick/doctron/pkg/alioss"
)

const (
	DoctronUploaderAliOss = "alioss"
	DoctronUploaderMock   = "mockUploader"
)

var LoadedConfig *Config
var OssConfig alioss.OssConfig

type Config struct {
	Doctron Doctron
	Oss     Oss
}

type Doctron struct {
	MaxConvertWorker int
	Env              string
	Retry            bool
	MaxConvertQueue  int
	ConvertTimeout   int
	Uploader         string
	Domain           string
	TLSCertFile      string
	TLSKeyFile       string
	User             []User
}

type User struct {
	Username string
	Password string
}

type Oss struct {
	Endpoint            string
	AccessKeyId         string
	AccessKeySecret     string
	BucketName          string
	PrivateServerDomain string
}

const defaultMaxConvertWorker = 50
const EnvProd = "prod"
const defaultRetry = false
const defaultMaxConvertQueue = 60
const defaultConvertTimeout = 30
const defaultUploader = DoctronUploaderAliOss
const mockUploader = DoctronUploaderMock
const defaultDomain = ":8080"
const defaultUsername = "doctron"
const defaultPassword = "lampnick"

func NewConfig() *Config {
	return &Config{
		Doctron: Doctron{
			MaxConvertWorker: defaultMaxConvertWorker,
			Env:              EnvProd,
			Retry:            defaultRetry,
			MaxConvertQueue:  defaultMaxConvertQueue,
			ConvertTimeout:   defaultConvertTimeout,
			Uploader:         defaultUploader,
			Domain:           defaultDomain,
			TLSCertFile:      "",
			TLSKeyFile:       "",
			User: []User{
				{
					Username: defaultUsername,
					Password: defaultPassword,
				},
			},
		},
		Oss: Oss{},
	}
}

func NewMockConfig() *Config {
	config := NewConfig()
	config.Doctron.Uploader = mockUploader
	return config
}

func (conf *Config) String() string {
	b, err := json.Marshal(*conf)
	if err != nil {
		return fmt.Sprintf("%+v", *conf)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *conf)
	}
	return out.String()
}
