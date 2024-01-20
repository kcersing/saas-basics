package db

import (
	"context"
	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"saas_basics/conf"
	"saas_basics/db/ent"
	"saas_basics/db/ent/migrate"
	"time"
)

var DB *ent.Client

// Init init DB
func Init() {
	var err error
	mysqlConfig := conf.Conf().MySql.Name + ":" + conf.Conf().MySql.Password + "@tcp(" + conf.Conf().MySql.Url + ":" + conf.Conf().MySql.Port + ")/" + conf.Conf().MySql.Database + "?parseTime=True"

	drvWd, err := sql.Open("mysql", mysqlConfig)
	if err != nil {
		panic(err)
	}
	db := drvWd.DB()
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(1000)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(time.Minute * 5)

	DB = ent.NewClient(ent.Driver(drvWd))
	ctx := context.Background()
	if err := DB.Schema.Create(
		ctx,
		migrate.WithGlobalUniqueID(false),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
