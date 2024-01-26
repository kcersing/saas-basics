package config

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Key  string `mapstructure:"key" json:"key"`
}

type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}

type WXConfig struct {
	AppId     string `mapstructure:"app_id" json:"app_id"`
	AppSecret string `mapstructure:"app_secret" json:"app_secret"`
}

type PasetoConfig struct {
	SecretKey string `mapstructure:"secret_key" json:"secret_key"`
	Implicit  string `mapstructure:"implicit" json:"implicit"`
}

type ServerConfig struct {
	Name       string       `mapstructure:"name" json:"name"`
	Host       string       `mapstructure:"host" json:"host"`
	PasetoInfo PasetoConfig `mapstructure:"paseto" json:"paseto"`
	OtelInfo   OtelConfig   `mapstructure:"otel" json:"otel"`
	WXInfo     WXConfig     `mapstructure:"wx_config" json:"wx_config"`
	MysqlInfo  MysqlConfig  `mapstructure:"mysql" json:"mysql"`
}

type MysqlConfig struct {
	Source string `mapstructure:"source" json:"source"`
	Salt   string `mapstructure:"salt" json:"salt"`
}
