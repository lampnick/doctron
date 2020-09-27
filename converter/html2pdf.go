package converter

import (
	"context"
	"errors"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

// PrintToPDFParams print page as PDF.
type PDFParams struct {
	Params page.PrintToPDFParams
}

type html2pdf struct {
	buf []byte
}

func NewHtml2pdf() *html2pdf {
	return &html2pdf{}
}

func (ins *html2pdf) Convert(ctx context.Context, cc ConvertConfig) ([]byte, error) {
	var params PDFParams
	params, ok := cc.Params.(PDFParams)
	if !ok {
		return nil, errors.New("wrong pdf params given")
	}
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	if err := chromedp.Run(ctx,
		chromedp.Navigate(cc.Url),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			ins.buf, _, err = params.Params.Do(ctx)
			return err
		}),
	); err != nil {
		return nil, err
	}

	return ins.buf, nil
}

func (ins *html2pdf) Upload(ctx context.Context, co ConvertOutput) (string, error) {

	return "", nil
}
