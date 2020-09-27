package doctron_core

import (
	"context"
	"errors"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"time"
)

// Image compression format (defaults to png).
const FormatPng = page.CaptureScreenshotFormatPng
const FormatJpeg = page.CaptureScreenshotFormatJpeg

// Compression quality from range [0..100] (jpeg only).
const DefaultQuality = 0

// Capture the screenshot from the surface, rather than the view. Defaults to true.
const DefaultFromSurface = true

// Capture the screenshot of a given region only.
// X offset in device independent pixels (dip).
const DefaultViewportX = 0

// Y offset in device independent pixels (dip).
const DefaultViewportY = 0

// Rectangle width in device independent pixels (dip).
const DefaultViewportWidth = 996

// Rectangle height in device independent pixels (dip).
const DefaultViewportHeight = 996

// Page scale factor.
const DefaultViewportScale = 1

// PrintToHtml2Image print page as PDF.
type Html2ImageParams struct {
	page.CaptureScreenshotParams
	CustomClip bool
}

func NewDefaultHtml2ImageParams() Html2ImageParams {
	return Html2ImageParams{
		CustomClip: false,
		CaptureScreenshotParams: page.CaptureScreenshotParams{
			Format:  FormatPng,
			Quality: DefaultQuality,
			Clip: &page.Viewport{
				X:      DefaultViewportX,
				Y:      DefaultViewportY,
				Width:  DefaultViewportWidth,
				Height: DefaultViewportHeight,
				Scale:  DefaultViewportScale,
			},
			FromSurface: DefaultFromSurface,
		},
	}
}

type html2image struct {
	Doctron
}

func (ins *html2image) GetConvertElapsed() time.Duration {
	return ins.convertElapsed
}

func (ins *html2image) Convert() ([]byte, error) {
	start := time.Now()
	defer func() {
		ins.convertElapsed = time.Since(start)
	}()
	var params Html2ImageParams
	params, ok := ins.cc.Params.(Html2ImageParams)
	if !ok {
		return nil, errors.New("wrong html2image params given")
	}
	ctx, cancel := chromedp.NewContext(ins.ctx)
	defer cancel()

	if err := chromedp.Run(ctx,
		chromedp.Navigate(ins.cc.Url),
		chromedp.ActionFunc(func(ctx context.Context) error {

			if !params.CustomClip {
				// get layout metrics
				_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
				if err != nil {
					return err
				}
				params.Clip.X = contentSize.X
				params.Clip.Y = contentSize.Y
				params.Clip.Width = contentSize.Width
				params.Clip.Height = contentSize.Height
			}

			//width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// force viewport emulation
			err := emulation.SetDeviceMetricsOverride(int64(params.Clip.Width), int64(params.Clip.Height), 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			// capture screenshot
			ins.buf, err = params.Do(ctx)
			return err
		}),
	); err != nil {
		return nil, err
	}

	return ins.buf, nil
}
