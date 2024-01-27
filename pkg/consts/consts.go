package consts

import "time"

const (
	FreeCar    = "FreeCar"
	Issuer     = "FreeCar"
	Admin      = "Admin"
	User       = "User"
	ThirtyDays = time.Hour * 24 * 30
	AccountID  = "accountID"
	ID         = "id"

	ApiConfigPath   = "./app/api/config.yaml"
	ApiConfigPath1  = "./app/api/tsconfig.json"
	UserConfigPath  = "./app/user/config.yaml"
	UserConfigPath1 = "./app/user/tsconfig.json"

	HlogFilePath = "./tmp/hlog/logs/"
	KlogFilePath = "./tmp/klog/logs/"

	MySqlDSN    = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	MongoURI    = "mongodb://%s:%s@%s:%d"
	RabbitMqURI = "amqp://%s:%s@%s:%d/"

	IPFlagName  = "ip"
	IPFlagValue = "0.0.0.0"
	IPFlagUsage = "address"

	PortFlagName  = "port"
	PortFlagUsage = "port"

	TCP = "tcp"

	FreePortAddress = "localhost:0"
	CorsAddress     = "http://localhost:3000"

	ConsulCheckInterval                       = "7s"
	ConsulCheckTimeout                        = "5s"
	ConsulCheckDeregisterCriticalServiceAfter = "15s"

	CarCollection     = "car"
	ProfileCollection = "profile"
	TripCollection    = "trip"

	RedisProfileClientDB = 0
	RedisBlobClientDB    = 0
	RedisCarClientDB     = 0

	UserSnowflakeNode  = 2
	BlobSnowflakeNode  = 3
	AdminSnowflakeNode = 4

	LimitOfSomeCars     = 20
	LimitOfSomeTrips    = 20
	LimitOfSomeProfiles = 20
	LimitOfSomeUsers    = 20

	OCRUrl = "https://aip.baidubce.com/rest/2.0/ocr/v1/driving_license"

	GPTUrl = "https://api.openai.com/v1/chat/completions"
)
