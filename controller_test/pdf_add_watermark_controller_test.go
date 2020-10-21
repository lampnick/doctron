package controller_test

import (
	"github.com/Jeffail/tunny"
	"github.com/kataras/iris/v12/httptest"
	"github.com/lampnick/doctron/app"
	"github.com/lampnick/doctron/conf"
	"github.com/lampnick/doctron/mock"
	"github.com/lampnick/doctron/worker"
	"io/ioutil"
	"testing"
)

func init() {
	conf.LoadedConfig = conf.NewConfig()
	conf.LoadedConfig.Doctron.Uploader = conf.DoctronUploaderMock
	conf.LoadedConfig.Oss.PrivateServerDomain = "www.lampnick.com"
	worker.Pool = tunny.NewFunc(conf.LoadedConfig.Doctron.MaxConvertWorker, worker.DoctronHandler)
}

func TestPdfAddWatermark(t *testing.T) {
	loadedPdf, err := ioutil.ReadFile("../test_data/doctron.pdf")
	if err != nil {
		t.Errorf("loaded pdf failed,%s", err.Error())
	}
	ts := mock.HTTPServerByte("application/pdf", loadedPdf, false)
	defer ts.Close()

	loadedImage, err := ioutil.ReadFile("../test_data/doctron.png")
	if err != nil {
		t.Errorf("loaded png failed,%s", err.Error())
	}
	tsImg := mock.HTTPServerByte("image/png", loadedImage, false)
	defer tsImg.Close()

	doctron := app.NewDoctron()
	expect := httptest.New(t, doctron)
	request := expect.GET("/convert/pdfAddWatermark")
	request.WithQuery("u", "doctron")
	request.WithQuery("p", "lampnick")
	request.WithQuery("imageUrl", tsImg.URL)
	request.WithQuery("url", ts.URL)
	response := request.Expect().Status(httptest.StatusOK)
	response.Body().Length().EqualDelta(4191, 4193)
}

func TestPdfAddWatermarkUpload(t *testing.T) {
	loadedPdf, err := ioutil.ReadFile("../test_data/doctron.pdf")
	if err != nil {
		t.Errorf("loaded pdf failed,%s", err.Error())
	}
	ts := mock.HTTPServerByte("application/pdf", loadedPdf, false)
	defer ts.Close()

	loadedImage, err := ioutil.ReadFile("../test_data/doctron.png")
	if err != nil {
		t.Errorf("loaded png failed,%s", err.Error())
	}
	tsImg := mock.HTTPServerByte("image/png", loadedImage, false)
	defer tsImg.Close()

	doctron := app.NewDoctron()
	expect := httptest.New(t, doctron)
	request := expect.GET("/convert/pdfAddWatermark")
	request.WithQuery("u", "doctron")
	request.WithQuery("p", "lampnick")
	request.WithQuery("imageUrl", tsImg.URL)
	request.WithQuery("url", ts.URL)
	request.WithQuery("uploadKey", "doctron.pdf")
	response := request.Expect().Status(httptest.StatusOK)
	expected := `{"code":0,"message":"","data":"http://www.lampnick.com/doctron.pdf"}`
	response.Body().Equal(expected)
}
