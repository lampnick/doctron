package doctron_core

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/lampnick/doctron/pkg/curl"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type PdfWatermarkParams struct {
	WatermarkType int
	ImageUrl      string
}

func NewDefaultPdfWatermarkParams() PdfWatermarkParams {
	return PdfWatermarkParams{
		WatermarkType: 0,
		ImageUrl:      "",
	}
}

type pdfWatermark struct {
	Doctron
}

func (ins pdfWatermark) Convert() ([]byte, error) {
	start := time.Now()
	defer func() {
		ins.convertElapsed = time.Since(start)
	}()
	var eg errgroup.Group
	var pdfReader *bytes.Reader
	var watermark string
	var watermarkedPdf *os.File
	defer func() {
		if watermarkedPdf != nil {
			err := os.Remove(watermarkedPdf.Name())
			if err != nil {
				log.Printf("remove watermarkedPdf temp file failed,reason: " + err.Error())
			}
		}
		if _, err := os.Stat(watermark); err == nil {
			err := os.Remove(watermark)
			if err != nil {
				log.Printf("remove watermark temp file failed,reason: " + err.Error())
			}
		}
	}()
	var params PdfWatermarkParams
	params, ok := ins.cc.Params.(PdfWatermarkParams)
	if !ok {
		return nil, errors.New("wrong pdf watermark params given")
	}
	eg.Go(func() (err error) {
		defer func() {
			e := recover()
			if e != nil {
				err = fmt.Errorf("recover err:%v", e)
			}
		}()
		pdfBytes, err := curl.GetBytesFromUrl(ins.cc.Url)
		if err != nil {
			return errors.New("download pdf failed,reason: " + err.Error())
		}
		if len(pdfBytes) == 0 {
			return errors.New("empty pdf given,please check pdf url")
		}
		pdfReader = bytes.NewReader(pdfBytes)
		return
	})
	eg.Go(func() (err error) {
		defer func() {
			e := recover()
			if e != nil {
				err = fmt.Errorf("recover err:%v", e)
			}
		}()
		watermarkBytes, err := curl.GetBytesFromUrl(params.ImageUrl)
		if err != nil {
			return err
		}
		if len(watermarkBytes) == 0 {
			return errors.New("empty image given,please check image url")
		}
		watermark = os.TempDir() + fmt.Sprintf("%s.png", uuid.NewV4())
		err = ioutil.WriteFile(watermark, watermarkBytes, 0777)
		if err != nil {
			return errors.New("write to image temp file failed,reason: " + err.Error())
		}
		return
	})
	eg.Go(func() (err error) {
		defer func() {
			e := recover()
			if e != nil {
				err = fmt.Errorf("recover err:%v", e)
			}
		}()
		watermarkedPdf, err = ioutil.TempFile(os.TempDir(), fmt.Sprintf("%s.pdf", uuid.NewV4()))
		return
	})
	if err := eg.Wait(); err != nil {
		return nil, err
	}

	wm, err := pdfcpu.ParseImageWatermarkDetails(watermark, "", true)
	if err != nil {
		return []byte{}, errors.New("parse image failed,reason: " + err.Error())
	}

	err = api.AddWatermarks(pdfReader, watermarkedPdf, nil, wm, nil)
	if err != nil {
		return []byte{}, errors.New("add watermark to pdf failed,reason: " + err.Error())
	}
	err = watermarkedPdf.Sync()
	if err != nil {
		return []byte{}, errors.New("sync watermarked pdf to disk failed,reason: " + err.Error())
	}
	err = watermarkedPdf.Close()
	if err != nil {
		return []byte{}, errors.New("close watermarked pdf failed,reason: " + err.Error())
	}

	return ioutil.ReadFile(watermarkedPdf.Name())

}

func (ins pdfWatermark) GetConvertElapsed() time.Duration {
	return ins.convertElapsed
}
