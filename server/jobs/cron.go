package jobs

import (
	"github.com/flipped-aurora/gin-vue-admin/server/jobs/cron_task"
	"github.com/robfig/cron/v3"
)

func CronInit() {
	c := cron.New(cron.WithSeconds())
	refRefreshAccessToken, _ := c.AddFunc("*/10 * * * * *", cron_task.RefreshAccessToken)             // 每10s刷新一次微信access token
	refRefreshPublicAccessToken, _ := c.AddFunc("*/10 * * * * *", cron_task.RefreshPublicAccessToken) // 每10s刷新一次微信公众号access token
	refRefreshPublicJsApiTicket, _ := c.AddFunc("*/10 * * * * *", cron_task.RefreshPublicJsApiTicket) // 每10s刷新一次微信公众号js api ticket

	c.Start()

	c.Entry(refRefreshAccessToken).Job.Run()       // 立即执行一次
	c.Entry(refRefreshPublicAccessToken).Job.Run() // 立即执行一次
	c.Entry(refRefreshPublicJsApiTicket).Job.Run() // 立即执行一次
}
