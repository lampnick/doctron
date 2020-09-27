package doctron_core

import (
	"context"
	"errors"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"time"
)

// Paper orientation. Defaults to false.
const DefaultLandscape = false

// Display header and footer. Defaults to false.
const DefaultDisplayHeaderFooter = false

// Print background graphics. Defaults to true.
const DefaultPrintBackground = true

// Scale of the webpage rendering. Defaults to 1.
const DefaultScale = 1

// Paper width in inches. Defaults to 8.5 inches.
const DefaultPaperWidth = 8.5

// Paper height in inches. Defaults to 11 inches.
const DefaultPaperHeight = 11

// Top margin in inches. Defaults to 1cm (~0.4 inches).
const DefaultMarginTop = 0.4

// Bottom margin in inches. Defaults to 1cm (~0.4 inches).
const DefaultMarginBottom = 0.4

// Left margin in inches. Defaults to 1cm (~0.4 inches).
const DefaultMarginLeft = 0.4

// Right margin in inches. Defaults to 1cm (~0.4 inches).
const DefaultMarginRight = 0.4

// Paper ranges to print, e.g., '1-5, 8, 11-13'. Defaults to the empty string, which means print all pages.
const DefaultPageRanges = ""

// Whether to silently ignore invalid but successfully parsed page ranges, such as '3-2'. Defaults to false.
const DefaultIgnoreInvalidPageRanges = false

// HTML template for the print header. Should be valid HTML markup with following classes used to inject printing values into them: - date: formatted print date - title: document title - url: document location - pageNumber: current page number - totalPages: total pages in the document  For example, <span class=title></span> would generate span containing the title.
const DefaultHeaderTemplate = ""

// HTML template for the print footer. Should use the same format as the headerTemplate.
const DefaultFooterTemplate = ""

// Whether or not to prefer page size as defined by css. Defaults to false, in which case the content will be scaled to fit the paper size.
const DefaultPreferCSSPageSize = false

// PrintToPDFParams print page as PDF.
type PDFParams struct {
	page.PrintToPDFParams
}

func NewDefaultPDFParams() PDFParams {
	return PDFParams{
		page.PrintToPDFParams{
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
