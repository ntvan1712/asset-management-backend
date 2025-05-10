package infras

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"sync"
	"time"

	"asset_management_backend/app_config"
	"asset_management_backend/common/error_app"
	"asset_management_backend/common/logger"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var (
	dbProvider *DbProvider
	initDbOnce sync.Once
)

type DbProvider struct {
	Instance *bun.DB
	Config   app_config.PostgresConfig
}

func GetDbProvider() *DbProvider {
	initDbOnce.Do(initDb)
	return dbProvider
}

func initDb() {

	postgresqlConfig := app_config.GetAppConfig().PostgresConfig
	dsn := postgresqlConfig.ConnectionString

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	dbProvider = &DbProvider{
		Instance: bun.NewDB(sqldb, pgdialect.New(), bun.WithDiscardUnknownColumns()),
		Config:   postgresqlConfig,
	}
	dbProvider.Instance.AddQueryHook(&QueryHook{})

	logger.Info("[PostgresInfras] init PostgresDB với Bun")
}

func BuildUpdateQueryByMap(query *bun.UpdateQuery, updateData map[string]interface{}) {
	for key, value := range updateData {
		query = query.Set(fmt.Sprintf("%s = ?", key), value)
	}
}

func DeleteByID(ctx context.Context, db *bun.DB, tableName string, id int) error {
	res, err := db.NewDelete().
		Table(tableName).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return error_app.ErrDocumentNotFound
	}
	return nil
}

type QueryHook struct{}

func (h *QueryHook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	return ctx
}

func (h *QueryHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	query := string(event.Query)
	if strings.HasPrefix(query, "NOTIFY") {
		return // bỏ qua notify
	}
	logger.Info("[BunQueryHook]", time.Since(event.StartTime).String(), query)
}
