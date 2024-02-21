package cron_task

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
)

func recoverPanic() {
	if err := recover(); err != nil {
		global.GVA_LOG.Error("panic caught:", zap.Any("err", err))
	}
}
