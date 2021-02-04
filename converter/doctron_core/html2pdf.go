package doctron_core

import (
	"context"
	"errors"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

//DefaultLandscape Paper orientation. Defaults to false.
const DefaultLandscape = false

//DefaultDisplayHeaderFooter Display header and footer. Defaults to false.
const DefaultDisplayHeaderFooter = false

//DefaultPrintBackground Print background graphics. Defaults to true.
const DefaultPrintBackground = true

//DefaultScale Scale of the webpage rendering. Defaults to 1.
const DefaultScale = 1

//DefaultPaperWidth Paper width in inches. Defaults to 8.5 inches.
const DefaultPaperWidth = 8.5

//DefaultPaperHeight Paper height in inches. Defaults to 11 inches.
const DefaultPaperHeight = 11

//DefaultMarginTop Top margin in inches. Defaults to 1cm (~0.4 inches).
const DefaultMarginTop = 0.4

//DefaultMarginBottom Bottom margin in inches. Defaults to 1cm (~0.4 inches).
const DefaultMarginBottom = 0.4

//DefaultMarginLeft Left margin in inches. Defaults to 1cm (~0.4 inches).
const DefaultMarginLeft = 0.4

//DefaultMarginRight Right margin in inches. Defaults to 1cm (~0.4 inches).
const DefaultMarginRight = 0.4

//DefaultPageRanges Paper ranges to print, e.g., '1-5, 8, 11-13'. Defaults to the empty string, which means print all pages.
const DefaultPageRanges = ""

//DefaultIgnoreInvalidPageRanges Whether to silently ignore invalid but successfully parsed page ranges, such as '3-2'. Defaults to false.
const DefaultIgnoreInvalidPageRanges = false

//DefaultHeaderTemplate HTML template for the print header. Should be valid HTML markup with following classes used to inject printing values into them: - date: formatted print date - title: document title - url: document location - pageNumber: current page number - totalPages: total pages in the document  For example, <span class=title></span> would generate span containing the title.
const DefaultHeaderTemplate = ""

//DefaultFooterTemplate HTML template for the print footer. Should use the same format as the headerTemplate.
const DefaultFooterTemplate = ""

//DefaultPreferCSSPageSize Whether or not to prefer page size as defined by css. Defaults to false, in which case the content will be scaled to fit the paper size.
const DefaultPreferCSSPageSize = false

//DefaultWaitingTime Waiting time after the page loaded. Default 0 means not wait. unit:Millisecond
const DefaultWaitingTime = 0

// PDFParams print page as PDF.
type PDFParams struct {
	page.PrintToPDFParams
	WaitingTime int // Waiting time after the page loaded. Default 0 means not wait. unit:Millisecond
}

//NewDefaultPDFParams default pdf params
func NewDefaultPDFParams() PDFParams {
	return PDFParams{
		PrintToPDFParams: page.PrintToPDFParams{
			Landscape:               DefaultLandscape,
			DisplayHeaderFooter:     DefaultDisplayHeaderFooter,
			PrintBackground:         DefaultPrintBackground,
			Scale:                   DefaultScale,
			PaperWidth:              DefaultPaperWidth,
			PaperHeight:             DefaultPaperHeight,
			MarginTop:               DefaultMarginTop,
			MarginBottom:            DefaultMarginBottom,
			MarginLeft:              DefaultMarginLeft,
			MarginRight:             DefaultMarginRight,
			PageRanges:              DefaultPageRanges,
			IgnoreInvalidPageRanges: DefaultIgnoreInvalidPageRanges,
			HeaderTemplate:          DefaultHeaderTemplate,
			FooterTemplate:          DefaultFooterTemplate,
			PreferCSSPageSize:       DefaultPreferCSSPageSize,
		},
		WaitingTime: DefaultWaitingTime,
	}
}

type html2pdf struct {
	Doctron
}

func (ins *html2pdf) GetConvertElapsed() time.Duration {
	return ins.convertElapsed
}

func (ins *html2pdf) Convert() ([]byte, error) {
	start := time.Now()
	defer func() {
		ins.convertElapsed = time.Since(start)
	}()
	var params PDFParams
	params, ok := ins.cc.Params.(PDFParams)
	if !ok {
		return nil, errors.New("wrong pdf params given")
	}
	ctx, cancel := chromedp.NewContext(ins.ctx)
	defer cancel()

	if err := chromedp.Run(ctx,
		chromedp.Navigate(ins.cc.Url),
		chromedp.Sleep(time.Duration(params.WaitingTime)*time.Millisecond),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			ins.buf, _, err = params.Do(ctx)
			return err
		}),
	); err != nil {
		return nil, err
	}

	return ins.buf, nil
}
