package conf

import (
	"bytes"
	"encoding/json"
	"fmt"
)

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
	TLSCertFile	string
	TLSKeyFile string
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

func NewConfig() *Config {
	return &Config{}
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