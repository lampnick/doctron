package controller

import "github.com/chromedp/cdproto/page"

type CommonRequestDTO struct {
	Url       string `schema:"url,omitempty" validate:"required,url"`
	UploadKey string `schema:"uploadKey,omitempty" validate:"omitempty"`
	Username  string `schema:"u,omitempty" validate:"required"`
	Password  string `schema:"p,omitempty" validate:"required"`
}

type Html2PdfRequestDTO struct {
	CommonRequestDTO
	Landscape               bool    `schema:"landscape,omitempty" validate:"omitempty"`               // Paper orientation. core.Defaults to false.
	DisplayHeaderFooter     bool    `schema:"displayHeaderFooter,omitempty" validate:"omitempty"`     // Display header and footer. core.Defaults to false.
	PrintBackground         bool    `schema:"printBackground,omitempty" validate:"omitempty"`         // Print background graphics. core.Defaults to false.
	Scale                   float64 `schema:"scale,omitempty" validate:"omitempty"`                   // Scale of the webpage rendering. core.Defaults to 1.
	PaperWidth              float64 `schema:"paperWidth,omitempty" validate:"gt=0"`                   // Paper width in inches. core.Defaults to 8.5 inches.
	PaperHeight             float64 `schema:"paperHeight,omitempty" validate:"gt=0"`                  // Paper height in inches. core.Defaults to 11 inches.
	MarginTop               float64 `schema:"marginTop" validate:"gte=0"`                             // Top margin in inches. core.Defaults to 1cm (~0.4 inches).
	MarginBottom            float64 `schema:"marginBottom" validate:"gte=0"`                          // Bottom margin in inches. core.Defaults to 1cm (~0.4 inches).
	MarginLeft              float64 `schema:"marginLeft" validate:"gte=0"`                            // Left margin in inches. core.Defaults to 1cm (~0.4 inches).
	MarginRight             float64 `schema:"marginRight" validate:"gte=0"`                           // Right margin in inches. core.Defaults to 1cm (~0.4 inches).
	PageRanges              string  `schema:"pageRanges,omitempty" validate:"omitempty"`              // Paper ranges to print, e.g., '1-5, 8, 11-13'. core.Defaults to the empty string, which means print all pages.
	IgnoreInvalidPageRanges bool    `schema:"ignoreInvalidPageRanges,omitempty" validate:"omitempty"` // Whether to silently ignore invalid but successfully parsed page ranges, such as '3-2'. core.Defaults to false.
	HeaderTemplate          string  `schema:"headerTemplate,omitempty" validate:"omitempty"`          // HTML template for the print header. Should be valid HTML markup with following classes used to inject printing values into them: - date: formatted print date - title: document title - url: document location - pageNumber: current page number - totalPages: total pages in the document  For example, <span class=title></span> would generate span containing the title.
	FooterTemplate          string  `schema:"footerTemplate,omitempty" validate:"omitempty"`          // HTML template for the print footer. Should use the same format as the headerTemplate.
	PreferCSSPageSize       bool    `schema:"preferCSSPageSize,omitempty" validate:"omitempty"`       // Whether or not to prefer page size as defined by css. core.Defaults to false, in which case the content will be scaled to fit the paper size.
	WaitingTime             int     `schema:"waitingTime,omitempty" validate:"omitempty"`             // Waiting time after the page loaded. Default 0 means not wait. unit:Millisecond
}

type Html2ImageRequestDTO struct {
	CommonRequestDTO
	Format      page.CaptureScreenshotFormat `schema:"format,omitempty" validate:"omitempty"`      // Image compression format (defaults to png).
	Quality     int64                        `schema:"quality,omitempty" validate:"omitempty"`     // Compression quality from range [0..100] (jpeg only).
	CustomClip  bool                         `schema:"customClip,omitempty" validate:"omitempty"`  //if set this value, the below clip will work,otherwise not work!
	ClipX       float64                      `schema:"clipX,omitempty" validate:"omitempty"`       // Capture the screenshot of a given region only.X offset in device independent pixels (dip).
	ClipY       float64                      `schema:"clipY,omitempty" validate:"omitempty"`       // Capture the screenshot of a given region only.Y offset in device independent pixels (dip).
	ClipWidth   float64                      `schema:"clipWidth,omitempty" validate:"omitempty"`   // Capture the screenshot of a given region only.Rectangle width in device independent pixels (dip).
	ClipHeight  float64                      `schema:"clipHeight,omitempty" validate:"omitempty"`  // Capture the screenshot of a given region only.Rectangle height in device independent pixels (dip).
	ClipScale   float64                      `schema:"clipScale,omitempty" validate:"omitempty"`   // Capture the screenshot of a given region only.Page scale factor.
	FromSurface bool                         `schema:"fromSurface,omitempty" validate:"omitempty"` // Capture the screenshot from the surface, rather than the view. Defaults to true.
	WaitingTime int                          `schema:"waitingTime,omitempty" validate:"omitempty"` // Waiting time after the page loaded. Default 0 means not wait. unit:Millisecond
}

type PdfWatermarkRequestDTO struct {
	CommonRequestDTO
	WatermarkType int    `schema:"watermarkType,omitempty" validate:"omitempty"` // watermark type will support soon
	ImageUrl      string `schema:"imageUrl,omitempty" validate:"required,url"`   // watermark image url
}
