package config

type Wechat struct {
	AppId           string `mapstructure:"app-id" json:"appId" yaml:"app-id"`                                 // AppId
	AppSecret       string `mapstructure:"app-secret" json:"appSecret" yaml:"app-secret"`                     // 密钥
	PublicAppId     string `mapstructure:"public-app-id" json:"publicAppId" yaml:"public-app-id"`             // 公众号AppId
	PublicAppSecret string `mapstructure:"public-app-secret" json:"PublicAppSecret" yaml:"public-app-secret"` // 公众号AppId
}
