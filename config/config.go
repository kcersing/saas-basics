package config

type ServerConfig struct {
	Name           string           `mapstructure:"Name" json:"Name"`
	Host           string           `mapstructure:"Host" json:"Host"`
	Port           int              `mapstructure:"Port" json:"Port"`
	Timeout        int              `mapstructure:"Timeout" json:"Timeout"`
	IsProd         bool             `mapstructure:"IsProd" json:"IsProd"`
	MySQLInfo      MySQLConfig      `mapstructure:"MySQL" json:"MySQL"`
	PostgreSQLInfo PostgreSQLConfig `mapstructure:"PostgreSQL" json:"PostgreSQL"`
	Captcha        Captcha          `mapstructure:"Captcha" json:"Captcha"`
	Auth           Auth             `mapstructure:"Auth" json:"Auth"`
	Redis          Redis            `mapstructure:"Redis" json:"Redis"`
	Casbin         CasbinConf       `mapstructure:"Casbin" json:"Casbin"`
	Payment        Payment          `mapstructure:"Payment" json:"Payment"`
	Minio          Minio            `mapstructure:"Minio" json:"Minio"`
	Swagger        Swagger          `mapstructure:"Swagger" json:"Swagger"`
}

type MySQLConfig struct {
	Host string `mapstructure:"Host" json:"Host"`
	Salt string `mapstructure:"Salt" json:"Salt"`
}
type PostgreSQLConfig struct {
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
	Appid             string `mapstructure:"ModelText" yaml:"appid"`
	MchId             string `mapstructure:"ModelText" yaml:"mch_id"`
	ApiKey            string `mapstructure:"ModelText" yaml:"api_key"`
	ApiV3Key          string `mapstructure:"ModelText" yaml:"api_v3_key"`
	CertFileContent   string `mapstructure:"ModelText" yaml:"cert_file_content"`
	KeyFileContent    string `mapstructure:"ModelText" yaml:"key_file_content"`
	Pkcs12FileContent string `mapstructure:"ModelText" yaml:"pkcs12_file_content"`
}

type AliPay struct {
	Appid                   string `mapstructure:"ModelText" yaml:"appid"`
	PrivateKey              string `mapstructure:"ModelText" yaml:"private_key"`
	AppPublicCertContent    string `mapstructure:"ModelText" yaml:"app_public_cert_content"`
	AlipayRootCertContent   string `mapstructure:"ModelText" yaml:"alipay_root_cert_content"`
	AlipayPublicCertContent string `mapstructure:"ModelText" yaml:"alipay_public_cert_content"`
}

type Minio struct {
	EndPoint        string `mapstructure:"EndPoint" yaml:"EndPoint"`
	AccessKeyID     string `mapstructure:"AccessKeyID" yaml:"AccessKeyID"`
	SecretAccessKey string `mapstructure:"SecretAccessKey" yaml:"SecretAccessKey"`
	UseSSL          bool   `mapstructure:"UseSSL" yaml:"UseSSL"`

	VideoBucketName string `mapstructure:"VideoBucketName" yaml:"VideoBucketName"`
	ImgBucketName   string `mapstructure:"ImgBucketName" yaml:"ImgBucketName"`

	Url string `mapstructure:"Url" yaml:"Url"`
}
type Swagger struct {
	Url string `mapstructure:"Url" yaml:"Url"`
}
