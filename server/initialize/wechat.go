package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/tools"
)

func WechatInit() {
	appid := global.GVA_CONFIG.Wechat.AppId
	secret := global.GVA_CONFIG.Wechat.AppSecret
	publicAppId := global.GVA_CONFIG.Wechat.PublicAppId
	publicSecret := global.GVA_CONFIG.Wechat.PublicAppSecret
	global.WechatConfig = tools.NewWxConfig(appid, secret, publicAppId, publicSecret)
}
