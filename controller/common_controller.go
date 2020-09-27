package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/lampnick/doctron/common"
	"github.com/lampnick/doctron/worker"
	"runtime"
)

// server status
func ServerStatus(ctx iris.Context) {
	serverStatus := common.ServerStatus{
		Version:    common.Version,
		Goroutines: runtime.NumGoroutine(),
		Workers: worker.Pool.GetSize(),
		Queue: worker.Pool.QueueLength(),

	}
	_, _ = common.NewJsonOutput(ctx, common.NewDefaultOutputDTO(serverStatus))
}

func log(ctx iris.Context, format string, args ...interface{}) {
	ctx.Application().Logger().Infof(format, args...)
}
