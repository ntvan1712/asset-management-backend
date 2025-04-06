package infras

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"asset_management_backend/app_config"
	"asset_management_backend/common/logger"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var (
	dbInstance *bun.DB
	initDbOnce sync.Once
)

func GetDbInstance() *bun.DB {
	initDbOnce.Do(initDb)
	return dbInstance
}

func initDb() {
	dsn := app_config.GetAppConfig().PostgresConfig.ConnectionString

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	dbInstance = bun.NewDB(sqldb, pgdialect.New(), bun.WithDiscardUnknownColumns())
	dbInstance.AddQueryHook(&QueryHook{})

	logger.Info("[PostgresInfras] init PostgresDB với Bun")
}

type QueryHook struct{}

func (h *QueryHook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	return ctx
}

func (h *QueryHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	logger.Info("[BunQueryHook]",time.Since(event.StartTime).String(), string(event.Query))
}