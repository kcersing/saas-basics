package init

func Init() {
	InitConfig()
	InitCache()
	InitDB()
	NewInitDatabase()
	InitCasbin()
	InitLogger()
}
