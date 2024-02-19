package config

type ServerConfig struct {
	Name    string `mapstructure:"name" json:"name"`
	Host    string `mapstructure:"host" json:"host"`
	Port    int    `mapstructure:"port" json:"port"`
	Timeout int    `yaml:"Timeout"`

	MysqlInfo MysqlConfig `mapstructure:"MySql" json:"MySql"`
	Captcha   Captcha     `yaml:"Captcha"`
	Auth      Auth        `yaml:"Auth"`
	Redis     Redis       `yaml:"Redis"`
	Casbin    CasbinConf  `yaml:"Casbin"`
}

type MysqlConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Salt string `mapstructure:"salt" json:"salt"`
}
type Captcha struct {
	KeyLong   int `yaml:"KeyLong"`
	ImgWidth  int `yaml:"ImgWidth"`
	ImgHeight int `yaml:"ImgHeight"`
}
type Auth struct {
	OAuthKey     string `yaml:"OAuthKey"`
	AccessSecret string `yaml:"AccessSecret"`
	AccessExpire int    `yaml:"AccessExpire"`
}

type Redis struct {
	Host string `yaml:"Host"`
	Port int    `yaml:"Port"`
	Type string `yaml:"Type"`
}

type CasbinConf struct {
	ModelText string `yaml:"ModelText"`
}
