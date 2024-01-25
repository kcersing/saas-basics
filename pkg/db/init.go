package db

import (
	"context"
	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"saas/pkg/db/ent"
	"saas/pkg/db/ent/migrate"
	"time"
)

// InitDB init DB
func InitDB() *ent.Client {
	var err error
	mysqlConfig := "root:pass@tcp(localhost:3306)/test?parseTime=True"

	drvWd, err := sql.Open("mysql", mysqlConfig)
	if err != nil {
		panic(err)
	}
	db := drvWd.DB()
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(1000)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(time.Minute * 5)

	DB := ent.NewClient(ent.Driver(drvWd))
	ctx := context.Background()
	if err := DB.Schema.Create(
		ctx,
		migrate.WithGlobalUniqueID(false),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return DB
}
