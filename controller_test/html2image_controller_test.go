package controller_test

import (
	"github.com/Jeffail/tunny"
	"github.com/kataras/iris/v12/httptest"
	"github.com/lampnick/doctron/app"
	"github.com/lampnick/doctron/conf"
	"github.com/lampnick/doctron/mock"
	"github.com/lampnick/doctron/worker"
	"testing"
)

func init() {
	conf.LoadedConfig = conf.NewConfig()
	conf.LoadedConfig.Doctron.Uploader = conf.DoctronUploaderMock
	conf.LoadedConfig.Oss.PrivateServerDomain = "www.lampnick.com"
	worker.Pool = tunny.NewFunc(conf.LoadedConfig.Doctron.MaxConvertWorker,worker.DoctronHandler)
}

func TestHtml2Image(t *testing.T) {
	ts := mock.HTTPServer("text/html", "lampnick content test", false)
	defer ts.Close()

	doctron := app.NewDoctron()
	expect := httptest.New(t, doctron)
	request := expect.GET("/convert/html2image")
	request.WithQuery("u", "doctron")
	request.WithQuery("p", "lampnick")
	request.WithQuery("url", ts.URL)
	response := request.Expect().Status(httptest.StatusOK)
	response.Body().Length().Equal(6249)
}

func TestHtml2ImageUpload(t *testing.T) {
	ts := mock.HTTPServer("text/html", "lampnick content test", false)
	defer ts.Close()

	doctron := app.NewDoctron()
	expect := httptest.New(t, doctron)
	request := expect.GET("/convert/html2image")
	request.WithQuery("u", "doctron")
	request.WithQuery("p", "lampnick")
	request.WithQuery("url", ts.URL)
	request.WithQuery("uploadKey", "doctron.png")
	response := request.Expect().Status(httptest.StatusOK)
	expected := `{"code":0,"message":"","data":"http://www.lampnick.com/doctron.png"}`
	response.Body().Equal(expected)
}
