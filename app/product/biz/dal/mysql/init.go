package mysql

import (
	db2 "common/db"
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
	"log"
	"product/biz/dal/mysql/ent"
	"product/biz/dal/mysql/ent/migrate"
	"product/conf"
	"time"
)

var (
	DB  *ent.Client
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
	//if isProd {
	//	hlog.Info("prod mode, use default mysql driver")
	//	drive = drv
	//} else {
	// Debug driver.
	hlog.Info("dev mode, use debug mysql driver")
	drive = &db2.DebugTimeDriver{
		Driver: drv,
		Log: func(ctx context.Context, info ...any) {
			hlog.Info(info...)
		},
	}
	//}
	DB = ent.NewClient(ent.Driver(drive))

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
