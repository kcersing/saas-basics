package config

type ServerConfig struct {
	Name      string      `mapstructure:"name" json:"name"`
	Host      string      `mapstructure:"host" json:"host"`
	Port      int         `mapstructure:"port" json:"port"`
	MysqlInfo MysqlConfig `mapstructure:"MySql" json:"MySql"`
}

type MysqlConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Salt string `mapstructure:"salt" json:"salt"`
}
