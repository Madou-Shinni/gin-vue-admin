package common

import "fmt"

const (
	WeChatAccessTokenKey = "we_chat_access_token:%s"
	WeChatJsApiTicketKey = "we_chat_js_api_ticket:%s"
)

// GetWeChatAccessTokenKey 获取微信AccessToken的key
func GetWeChatAccessTokenKey(appid string) string {
	return fmt.Sprintf(WeChatAccessTokenKey, appid)
}

// GetWeChatJsApiTicketKey 获取微信JsApiTicket的key
func GetWeChatJsApiTicketKey(appid string) string {
	return fmt.Sprintf(WeChatJsApiTicketKey, appid)
}
