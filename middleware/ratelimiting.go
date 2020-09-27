package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/lampnick/doctron/common"
	"github.com/lampnick/doctron/conf"
	"github.com/lampnick/doctron/worker"
	"sync"
)

func CheckRateLimiting(ctx iris.Context) {
	mu := sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()
	if worker.Pool.QueueLength() > int64(conf.LoadedConfig.Doctron.MaxConvertQueue) {
		outputDTO := common.NewDefaultOutputDTO(nil)
		outputDTO.Code = common.ApiRateLimitExceeded
		_, _ = common.NewJsonOutput(ctx, outputDTO)
		return
	}
	ctx.Next()
}
