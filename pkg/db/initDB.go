package db

import (
	"context"
	"fmt"
	"log"
	migrate2 "saas/pkg/db/ent/migrate"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	_ "github.com/go-sql-driver/mysql"

	"saas/pkg/db/ent"
	  _ "github.com/lib/pq"
)

func OpenMySql(mysqlConfig string, isProd bool){
}


func OpenPq()(databaseUrl string) *ent.Client {
    db, err := sql.Open("pgx", databaseUrl)
    if err != nil {
        log.Fatal(err)
    }

    // Create an ent.Driver from `db`.
    drv := entsql.OpenDB(dialect.Postgres, db)
    return ent.NewClient(ent.Driver(drv))
}



// InitDB init DB
func InitDB(mysqlConfig string, isProd bool) (DB *ent.Client) {
	var err error
	drvWd, err := sql.Open("mysql", mysqlConfig)
	if err != nil {
		panic(err)
	}
	db := drvWd.DB()
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(1000)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(time.Minute * 5)

	var drive dialect.Driver
	if isProd {
		hlog.Info("prod mode, use default mysql driver")
		drive = drvWd
	} else {
		// Debug driver.
		hlog.Info("dev mode, use debug mysql driver")
		drive = &DebugTimeDriver{
			Driver: drvWd,
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

// DebugTimeDriver is a driver that logs all driver operations.
type DebugTimeDriver struct {
	dialect.Driver                               // underlying driver.
	log            func(context.Context, ...any) // log function. defaults to log.Println.
}

// Exec logs its params and calls the underlying driver Exec method.
func (d *DebugTimeDriver) Exec(ctx context.Context, query string, args, v any) error {
	start := time.Now()
	err := d.Driver.Exec(ctx, query, args, v)
	d.log(ctx, fmt.Sprintf("driver.Exec: query=%v args=%v time=%v", query, args, time.Since(start)))
	return err
}

// ExecContext logs its params and calls the underlying driver ExecContext method if it is supported.
func (d *DebugTimeDriver) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	drv, ok := d.Driver.(interface {
		ExecContext(context.Context, string, ...any) (sql.Result, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.ExecContext is not supported")
	}
	start := time.Now()
	result, err := drv.ExecContext(ctx, query, args...)
	d.log(ctx, fmt.Sprintf("driver.ExecContext: query=%v args=%v time=%v", query, args, time.Since(start)))
	return result, err
}

// Query logs its params and calls the underlying driver Query method.
func (d *DebugTimeDriver) Query(ctx context.Context, query string, args, v any) error {
	start := time.Now()
	err := d.Driver.Query(ctx, query, args, v)
	d.log(ctx, fmt.Sprintf("driver.Query: query=%v args=%v time=%v", query, args, time.Since(start)))
	return err
}

// QueryContext logs its params and calls the underlying driver QueryContext method if it is supported.
func (d *DebugTimeDriver) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	drv, ok := d.Driver.(interface {
		QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.QueryContext is not supported")
	}
	start := time.Now()
	rows, err := drv.QueryContext(ctx, query, args...)
	d.log(ctx, fmt.Sprintf("driver.QueryContext: query=%v args=%v time=%v", query, args, time.Since(start)))
	return rows, err
}
