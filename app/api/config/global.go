package config

import "saas/kitex_gen/cwg/user/userservice"

var (
	GlobalServerConfig ServerConfig
	GlobalConsulConfig ConsulConfig

	GlobalUserClient userservice.Client
)
