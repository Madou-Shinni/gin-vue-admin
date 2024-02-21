package tools

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
)

const (
	JsCode2SessionURL = "https://api.weixin.qq.com/sns/jscode2session"       // code2Session
	AccessTokenURL    = "https://api.weixin.qq.com/cgi-bin/stable_token"     // 获取access_token
	JsapiTicketURL    = "https://api.weixin.qq.com/cgi-bin/ticket/getticket" // 获取jsapi_ticket
	WebAccessTokenUrl = "https://api.weixin.qq.com/sns/oauth2/access_token"  // 获取特殊的网页授权access_token, 用于获取用户信息
	SnsapiUserInfoUrl = "https://api.weixin.qq.com/sns/userinfo"             // 获取用户信息
	GetQRCodeUrl      = "https://api.weixin.qq.com/wxa/getwxacode"           // 获取小程序码
)

type (
	// WxConfig 微信配置类
	WxConfig struct {
		AppID       string       `json:"appid"`         // 微信APPID
		Secret      string       `json:"secret"`        // 微信Secret
		PubWxConfig *PubWxConfig `json:"pub_wx_config"` // 公众号配置
	}

	// PubWxConfig 公众号配置类
	PubWxConfig struct {
		AppID  string `json:"appid"`  // 微信APPID
		Secret string `json:"secret"` // 微信Secret
	}

	// ErrorResp 错误返回
	ErrorResp struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}

	// GetJsCode2SessionResp 小程序解密密钥信息
	GetJsCode2SessionResp struct {
		Openid     string `json:"openid"`
		SessionKey string `json:"session_key"`
		UnionID    string `json:"unionid"`
		ErrCode    int    `json:"errcode"`
		ErrMsg     string `json:"errmsg"`
	}

	// GetStableAccessTokenResp 获取稳定的AccessToken
	GetStableAccessTokenResp struct {
		ErrCode     int    `json:"errcode"`
		ErrMsg      string `json:"errmsg"`
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}

	// GetJsapiTicketResp 获取用于调用微信JS接口的临时票据
	GetJsapiTicketResp struct {
		ErrCode   int    `json:"errcode"`
		ErrMsg    string `json:"errmsg"`
		Ticket    string `json:"ticket"`
		ExpiresIn int    `json:"expires_in"`
	}

	// GetJsApiUsingPermissions 获取用于调用微信JS接口的使用权限
	GetJsApiUsingPermissions struct {
		Noncestr    string `json:"noncestr"`
		JsapiTicket string `json:"jsapi_ticket"`
		Timestamp   int64  `json:"timestamp"`
		Url         string `json:"url"`
	}

	// GetJsApiUsingPermissionsResp 获取用于调用微信JS接口的使用权限返回值
	GetJsApiUsingPermissionsResp struct {
		GetJsApiUsingPermissions
		AppId     string `json:"appId"`
		Signature string `json:"signature"`
	}

	// GetWebAccessTokenResp 获取特殊的网页授权access_token, 用于获取用户信息
	GetWebAccessTokenResp struct {
		ErrorResp
		AccessToken    string `json:"access_token"`
		ExpiresIn      int    `json:"expires_in"`
		RefreshToken   string `json:"refresh_token"`
		OpenId         string `json:"openid"`
		Scope          string `json:"scope"`
		IsSnapshotuser bool   `json:"is_snapshotuser"`
		Unionid        string `json:"unionid"`
	}

	// GetSnsapiUserInfoResp 获取用户信息
	GetSnsapiUserInfoResp struct {
		ErrorResp
		Openid     string   `json:"openid"`
		Nickname   string   `json:"nickname"`
		Sex        int      `json:"sex"`
		Province   string   `json:"province"`
		City       string   `json:"city"`
		Country    string   `json:"country"`
		Headimgurl string   `json:"headimgurl"`
		Privilege  []string `json:"privilege"`
		Unionid    string   `json:"unionid"`
	}

	// GetPhoneNumberResp 获取用户手机号
	GetPhoneNumberResp struct {
		ErrorResp
		PhoneInfo struct {
			PhoneNumber     string `json:"phoneNumber"`
			PurePhoneNumber string `json:"purePhoneNumber"`
			CountryCode     string `json:"countryCode"`
			Watermark       struct {
				Appid     string `json:"appid"`
				Timestamp int64  `json:"timestamp"`
			}
		} `json:"phone_info"`
	}

	// GetQRCodeResp 获取小程序码错误返回
	GetQRCodeResp struct {
		ErrorResp
		Buffer []byte `json:"buffer"`
	}

	LineColor struct {
		R string `json:"r"`
		G string `json:"g"`
		B string `json:"b"`
	}
)

func NewWxConfig(appid string, secret string, publicAppId string, publicAppSecret string) *WxConfig {
	return &WxConfig{
		AppID:  appid,
		Secret: secret,
		PubWxConfig: &PubWxConfig{
			AppID:  publicAppId,
			Secret: publicAppSecret,
		},
	}
}

// GetJsCode2Session 获取
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-login/code2Session.html 小程序登录
func (m *WxConfig) GetJsCode2Session(code string) (result *GetJsCode2SessionResp, err error) {
	data := map[string]interface{}{
		"appid":      m.AppID,
		"secret":     m.Secret,
		"js_code":    code,
		"grant_type": "authorization_code",
	}

	resp, err := NewRequest(GET, JsCode2SessionURL, data, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, fmt.Errorf("GetJsCode2Session error: %s", result.ErrMsg)
	}

	return
}

// GetStableAccessToken 获取稳定的AccessToken
func (m *WxConfig) GetStableAccessToken(force_refresh bool) (result *GetStableAccessTokenResp, err error) {
	data := map[string]interface{}{
		"appid":         m.AppID,
		"secret":        m.Secret,
		"force_refresh": force_refresh, // 有效期内重复调用该接口不会更新  true为强制刷新
		"grant_type":    "client_credential",
	}

	resp, err := NewRequest(POST, AccessTokenURL, data, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, err
	}

	return
}

// GetStableAccessToken 获取稳定的AccessToken
func (m *PubWxConfig) GetStableAccessToken(force_refresh bool) (result *GetStableAccessTokenResp, err error) {
	data := map[string]interface{}{
		"appid":         m.AppID,
		"secret":        m.Secret,
		"force_refresh": force_refresh, // 有效期内重复调用该接口不会更新  true为强制刷新
		"grant_type":    "client_credential",
	}

	resp, err := NewRequest(POST, AccessTokenURL, data, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, err
	}

	return
}

// GetWebAccessToken 获取特殊的网页授权access_token, 用于获取用户信息
func (m *PubWxConfig) GetWebAccessToken(code string) (result *GetWebAccessTokenResp, err error) {
	data := map[string]interface{}{
		"appid":      m.AppID,
		"secret":     m.Secret,
		"code":       code,
		"grant_type": "authorization_code",
	}

	resp, err := NewRequest(GET, WebAccessTokenUrl, data, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, err
	}

	return
}

// GetJsapiTicket 公众号：获取用于调用微信JS接口的临时票据
func GetJsapiTicket(accessToken string) (result *GetJsapiTicketResp, err error) {
	data := map[string]interface{}{
		"access_token": accessToken,
		"type":         "jsapi",
	}

	resp, err := NewRequest(GET, JsapiTicketURL, data, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, err
	}

	return
}

// Sha1 对所有待签名参数按照字段名的ASCII 码从小到大排序（字典序）后，使用URL键值对的格式（即key1=value1&key2=value2…）拼接成字符串string1
func (g GetJsApiUsingPermissions) Sha1() string {
	type fieldInfo struct {
		Name  string
		Value interface{}
	}

	v := reflect.ValueOf(g)
	t := v.Type()

	// 存储字段信息的切片
	var fields []fieldInfo

	// 遍历结构体字段，将字段信息存储到切片中
	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Tag.Get("json")
		fieldValue := v.Field(i).Interface()
		fields = append(fields, fieldInfo{Name: fieldName, Value: fieldValue})
	}

	// 按照字段名的 ASCII 码从小到大排序
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].Name < fields[j].Name
	})

	// 返回sha1值
	var string1, sha1 string
	for _, field := range fields {
		string1 = fmt.Sprint(string1, "&", field.Name, "=", field.Value)
	}
	string1 = removePrefix(string1)
	sha1 = Sha1Encrypt(string1)

	return sha1
}

// GetSnsapiUserInfo 公众号：拉取用户信息(需scope为 snsapi_userinfo)
func GetSnsapiUserInfo(webAccessToken, openId string) (result *GetSnsapiUserInfoResp, err error) {
	data := map[string]interface{}{
		"access_token": webAccessToken,
		"openid":       openId,
		"lang":         "zh_CN",
	}

	resp, err := NewRequest(GET, SnsapiUserInfoUrl, data, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, err
	}

	return
}

// GetPhoneNumber 小程序：获取用户手机号
func GetPhoneNumber(accessToken, code string) (result *GetPhoneNumberResp, err error) {
	data := map[string]interface{}{
		"code": code,
	}

	url := getPhoneNumUrl(accessToken)

	resp, err := NewRequest(POST, url, data, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, err
	}

	return
}

// getPhoneNumUrl 这边的url特殊需要组装url
func getPhoneNumUrl(accessToken string) string {
	return fmt.Sprint(SnsapiUserInfoUrl, "?", "access_token", accessToken)
}

// GetQRCode 小程序：获取小程序码
func GetQRCode(accessToken, path string, query map[string]interface{}, width int, auto_color bool, line_color *LineColor, is_hyaline bool, env_version string) (result *GetQRCodeResp, err error) {
	errResp := &GetQRCodeResp{}
	url := getQrCodeUrl(accessToken)
	data := map[string]interface{}{
		"path": getQrPath(path, query),
	}

	if width != 0 {
		data["width"] = width
	}
	if auto_color {
		data["auto_color"] = auto_color
	}
	if line_color != nil {
		data["line_color"] = line_color
	}
	if is_hyaline {
		data["is_hyaline"] = is_hyaline
	}
	if env_version != "" {
		data["env_version"] = env_version
	}

	resp, err := NewRequest(POST, url, data, nil)
	if err != nil {
		return nil, err
	}

	// 如果调用成功，会直接返回图片二进制内容，如果请求失败，会返回 JSON 格式的数据。
	json.Unmarshal(resp, &errResp)
	if errResp.ErrCode != 0 {
		return errResp, fmt.Errorf("GetQRCode error: %s", errResp.ErrMsg)
	}
	errResp.Buffer = resp

	return errResp, nil
}

// getQrCodeUrl 这边的url特殊需要组装url
func getQrCodeUrl(accessToken string) string {
	return fmt.Sprint(GetQRCodeUrl, "?", "access_token", "=", accessToken)
}

// getQrPath path需要加上query参数
func getQrPath(path string, queries map[string]interface{}) string {
	var queryStr string
	for k, v := range queries {
		queryStr = fmt.Sprint(queryStr, "&", k, "=", v)
	}
	queryStr = removePrefix(queryStr)
	return fmt.Sprint(path, "?", queryStr)
}
