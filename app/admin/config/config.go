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
