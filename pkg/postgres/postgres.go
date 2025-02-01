package postgres

import (
	"UchetUsers/internal/configs"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"time"
)

func ClientPostgres(logger *logrus.Logger, cfg configs.Config) (*pgxpool.Pool, error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLmode)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		logger.WithError(err).Error("Ошибка парсинга конфугурации подключения к Postgres")
		return nil, err
	}

	config.MaxConns = 10
	config.MinConns = 2
	config.HealthCheckPeriod = 30 * time.Second

	conn, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		logger.WithError(err).Error("Ошибка подключения к Postgres")
		return nil, err
	}

	logger.Info("Успешное подключение к Postgres")
	return conn, nil
}
