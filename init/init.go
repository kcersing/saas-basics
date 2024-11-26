package init

func Init() {
	InitConfig()
	InitCache()
	NewInitDatabase()
	InitCasbin()
	InitLogger()
}
