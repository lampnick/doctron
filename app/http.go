package app

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/core/router"
	"github.com/lampnick/doctron/controller"
	"github.com/lampnick/doctron/doctron_context"
	"github.com/lampnick/doctron/middleware"
)

func NewDoctron() *iris.Application {
	app := iris.Default()
	app.ContextPool.Attach(func() iris.Context {
		return &doctron_context.DoctronContext{
			// If you use the embedded Context,
			// call the `context.NewContext` to create one:
			Context: context.NewContext(app),
		}
	})
	app.PartyFunc("/convert", func(convert router.Party) {
		convert.Use(middleware.AuthMiddleware)
		convert.Use(middleware.CheckRateLimiting)
		convert.Get("/html2pdf", controller.Html2PdfHandler)
		convert.Get("/html2image", controller.Html2ImageHandler)
		convert.Get("/pdf2image", controller.Pdf2ImageHandler)
		convert.Get("/pdfAddWatermark", controller.PdfAddWatermarkHandler)
	})

	app.Handle("GET", "/status", controller.ServerStatus)

	return app
}
