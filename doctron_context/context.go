package doctron_context

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

// Create your own custom Context, put any fields you'll need.
type DoctronContext struct {
	// Embed the `iris.Context` -
	// It's totally optional but you will need this if you
	// don't want to override all the context's methods!
	iris.Context
}

// Optionally: validate DoctronContext implements iris.Context on compile-time.
var _ iris.Context = &DoctronContext{}

func (ctx *DoctronContext) Do(handlers context.Handlers) {
	context.Do(ctx, handlers)
}

func (ctx *DoctronContext) Next() {
	context.Next(ctx)
}

func (ctx *DoctronContext) Write(rawBody []byte) (int, error) {
	return ctx.ResponseWriter().Write(rawBody)
}
