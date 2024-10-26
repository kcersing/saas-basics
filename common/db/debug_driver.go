package db

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"

	"fmt"
	"time"
)

// DebugTimeDriver is a driver that logs all driver operations.
type DebugTimeDriver struct {
	dialect.Driver                               // underlying driver.
	Log            func(context.Context, ...any) // log function. defaults to log.Println.
}

// Exec logs its params and calls the underlying driver Exec method.
func (d *DebugTimeDriver) Exec(ctx context.Context, query string, args, v any) error {
	start := time.Now()
	err := d.Driver.Exec(ctx, query, args, v)
	d.Log(ctx, fmt.Sprintf("driver.Exec: query=%v args=%v time=%v", query, args, time.Since(start)))
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
	d.Log(ctx, fmt.Sprintf("driver.ExecContext: query=%v args=%v time=%v", query, args, time.Since(start)))
	return result, err
}

// Query logs its params and calls the underlying driver Query method.
func (d *DebugTimeDriver) Query(ctx context.Context, query string, args, v any) error {
	start := time.Now()
	err := d.Driver.Query(ctx, query, args, v)
	d.Log(ctx, fmt.Sprintf("driver.Query: query=%v args=%v time=%v", query, args, time.Since(start)))
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
	d.Log(ctx, fmt.Sprintf("driver.QueryContext: query=%v args=%v time=%v", query, args, time.Since(start)))
	return rows, err
}
