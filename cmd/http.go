package cmd

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"runtime"
)

func startHttp() {

	app := iris.Default()
	app.PartyFunc("/convert", func(convert router.Party) {
		convert.Use(authMiddleware)
		convert.Get("/html2pdf", html2pdfHandler)
		convert.Get("/html2image", html2imageHandler)
		convert.Get("/pdf2image", pdf2imageHandler)
		convert.Get("/pdfAddWatermark", pdfAddWatermarkHandler)
	})

	app.Handle("GET", "/status", func(ctx iris.Context) {
		_, _ = ctx.JSON(iris.Map{
			"version": version,
			"goroutines":runtime.NumGoroutine(),
		})
	})

	// Listens and serves incoming http requests
	// on http://localhost:8080.
	err := app.Listen(":8080")
	if err != nil {
		app.Logger().Fatal("start doctron failed. %v", err)
	}
}

func html2pdfHandler(ctx iris.Context) {

}

func html2imageHandler(ctx iris.Context) {

}

func pdf2imageHandler(ctx iris.Context) {

}

func pdfAddWatermarkHandler(ctx iris.Context) {

}

func authMiddleware(ctx iris.Context) {
	authParam := ctx.URLParam("auth")
	ctx.Application().Logger().Infof("authParam: %s", authParam)
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
