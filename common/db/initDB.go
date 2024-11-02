package db

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
	"log"
	"saas/pkg/db/ent"
	migrate2 "saas/pkg/db/ent/migrate"
	"time"
)

func OpenMySql(databaseUrl string) *entsql.Driver {
	db, err := sql.Open("mysql", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	// 从db变量中构造一个ent.Driver对象。
	drv := entsql.OpenDB(dialect.MySQL, db)

	return drv
}

func OpenPq(databaseUrl string) *entsql.Driver {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)

	return drv
}

// InitDB init DB
func InitDB(databaseUrl string, isProd bool) (DB *ent.Client) {
	drv := OpenPq(databaseUrl)

	// 生产环境使用默认mysql驱动，开发环境使用debug驱动
	var drive dialect.Driver
	if isProd {
		hlog.Info("prod mode, use default mysql driver")
		drive = drv
	} else {
		// Debug driver.
		hlog.Info("dev mode, use debug mysql driver")
		drive = &DebugTimeDriver{
			Driver: drv,
			log: func(ctx context.Context, info ...any) {
				hlog.Info(info...)
			},
		}
	}
	DB = ent.NewClient(ent.Driver(drive))

	ctx := context.Background()
	if err := DB.Schema.Create(
		ctx,
		migrate2.WithGlobalUniqueID(false),
		migrate2.WithDropIndex(true),
		migrate2.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// versioned-migration 不用
	// Create a local migration directory.
	//migrationsPath := "./pkg/db/ent/migrations"
	//_, err = os.Stat(migrationsPath)
	//if err != nil {
	//	if !os.IsNotExist(err) {
	//		hlog.Fatalf("failed to stat migrations path: %v", err)
	//		return nil
	//	}
	//	// Create the directory if it doesn't exist.
	//	err = os.MkdirAll(migrationsPath, os.ModePerm)
	//	if err != nil {
	//		hlog.Fatalf("failed creating migrations path: %v", err)
	//		return nil
	//	}
	//}
	//dir, err := migrate.NewLocalDir(migrationsPath)
	//if err != nil {
	//	hlog.Fatalf("failed creating atlas migration directory: %v", err)
	//	return nil
	//}
	//// Write migration diff.
	//err = schema.Diff(ctx, schema.WithDir(dir), schema.WithForeignKeys(false))
	//if err != nil {
	//	hlog.Fatalf("failed creating schema resources: %v", err)
	//	return nil
	//}
	//// Run the auto migration tool.
	//if err := DB.Schema.Create(ctx, schema.WithForeignKeys(false)); err != nil {
	//	hlog.Fatalf("failed creating schema resources: %v", err)
	//	return nil
	//}

	return DB
}
