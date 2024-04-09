package config

type ServerConfig struct {
	Name      string      `mapstructure:"Name" json:"Name"`
	Host      string      `mapstructure:"Host" json:"Host"`
	Port      int         `mapstructure:"Port" json:"Port"`
	Timeout   int         `mapstructure:"Timeout" json:"Timeout"`
	IsProd    bool        `mapstructure:"IsProd" json:"IsProd"`
	MysqlInfo MysqlConfig `mapstructure:"MySql" json:"MySql"`
	Captcha   Captcha     `mapstructure:"Captcha" json:"Captcha"`
	Auth      Auth        `mapstructure:"Auth" json:"Auth"`
	Redis     Redis       `mapstructure:"Redis" json:"Redis"`
	Casbin    CasbinConf  `mapstructure:"Casbin" json:"Casbin"`
}

type MysqlConfig struct {
	Host string `mapstructure:"Host" json:"Host"`
	Salt string `mapstructure:"Salt" json:"Salt"`
}
type Captcha struct {
	KeyLong   int `mapstructure:"KeyLong" json:"KeyLong"`
	ImgWidth  int `mapstructure:"ImgWidth" json:"ImgWidth"`
	ImgHeight int `mapstructure:"ImgHeight" json:"ImgHeight"`
}
type Auth struct {
	OAuthKey     string `mapstructure:"OAuthKey" json:"OAuthKey"`
	AccessSecret string `mapstructure:"AccessSecret" json:"AccessSecret"`
	AccessExpire int    `mapstructure:"AccessExpire" json:"AccessExpire"`
}

type Redis struct {
	Host string `mapstructure:"Host" json:"Host"`
	Port int    `mapstructure:"Port" json:"Port"`
	Type string `mapstructure:"Type" json:"Type"`
}

type CasbinConf struct {
	ModelText string `mapstructure:"ModelText" json:"ModelText"`
}

type Payment struct {
	Wechat *WechatPay `mapstructure:"Wechat" yaml:"Wechat"`
	Alipay *AliPay    `mapstructure:"Alipay" yaml:"Alipay"`
}

type WechatPay struct {
	Appid             string `yaml:"appid"`
	MchId             string `yaml:"mch_id"`
	ApiKey            string `yaml:"api_key"`
	ApiV3Key          string `yaml:"api_v3_key"`
	CertFileContent   string `yaml:"cert_file_content"`
	KeyFileContent    string `yaml:"key_file_content"`
	Pkcs12FileContent string `yaml:"pkcs12_file_content"`
}

type AliPay struct {
	Appid                   string `yaml:"appid"`
	PrivateKey              string `yaml:"private_key"`
	AppPublicCertContent    string `yaml:"app_public_cert_content"`
	AlipayRootCertContent   string `yaml:"alipay_root_cert_content"`
	AlipayPublicCertContent string `yaml:"alipay_public_cert_content"`
}
