package mysql

import (
	db2 "common/db"
	"entgo.io/ent/dialect"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"log"
	"order/biz/ent"
	"order/biz/ent/migrate"
	"order/conf"
	"time"

	"gorm.io/gorm"

	"context"
	"database/sql"

	entsql "entgo.io/ent/dialect/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {

	db, err := sql.Open("mysql", conf.GetConf().MySQL.DSN)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	// 从db变量中构造一个ent.Driver对象。
	drv := entsql.OpenDB(dialect.MySQL, db)

	// 生产环境使用默认mysql驱动，开发环境使用debug驱动
	var drive dialect.Driver
	if isProd {
		hlog.Info("prod mode, use default mysql driver")
		drive = drv
	} else {
		// Debug driver.
		hlog.Info("dev mode, use debug mysql driver")
		drive = &db2.DebugTimeDriver{
			Driver: drv,
			Log: func(ctx context.Context, info ...any) {
				hlog.Info(info...)
			},
		}
	}
	DB := ent.NewClient(ent.Driver(drive))

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
