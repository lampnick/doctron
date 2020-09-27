package doctron_core

import (
	"context"
	"github.com/lampnick/doctron/converter"
)

const (
	DoctronHtml2Pdf     = 1
	DoctronHtml2Image   = 2
	DoctronPdf2Image    = 3
	DoctronPdfWatermark = 4
)

// new doctron
func NewDoctron(ctx context.Context, doctronType int, cc converter.ConvertConfig) DoctronI {
	switch doctronType {
	case DoctronHtml2Pdf:
		fac := &(html2PdfFactory{})
		return fac.createDoctron(ctx, cc)
	case DoctronHtml2Image:
		fac := &(html2ImageFactory{})
		return fac.createDoctron(ctx, cc)
	case DoctronPdf2Image:
		fac := &(pdf2ImageFactory{})
		return fac.createDoctron(ctx, cc)
	case DoctronPdfWatermark:
		fac := &(pdfWatermarkFactory{})
		return fac.createDoctron(ctx, cc)
	default:
		return nil
	}
}

type DoctronFactory interface {
	createDoctron(ctx context.Context, cc converter.ConvertConfig) DoctronI
}

type html2PdfFactory struct {
}

func (ins *html2PdfFactory) createDoctron(ctx context.Context, cc converter.ConvertConfig) DoctronI {
	return &html2pdf{
		Doctron: Doctron{
			ctx: ctx,
			cc:  cc,
		},
	}
}

type html2ImageFactory struct {
}

func (ins *html2ImageFactory) createDoctron(ctx context.Context, cc converter.ConvertConfig) DoctronI {
	return &html2image{
		Doctron: Doctron{
			ctx: ctx,
			cc:  cc,
		},
	}
}

type pdf2ImageFactory struct {
}

func (ins *pdf2ImageFactory) createDoctron(ctx context.Context, cc converter.ConvertConfig) DoctronI {
	return &pdf2Image{
		Doctron: Doctron{
			ctx: ctx,
			cc:  cc,
		},
	}
}

type pdfWatermarkFactory struct {
}

func (ins *pdfWatermarkFactory) createDoctron(ctx context.Context, cc converter.ConvertConfig) DoctronI {
	return &pdfWatermark{
		Doctron: Doctron{
			ctx: ctx,
			cc:  cc,
		},
	}
}
