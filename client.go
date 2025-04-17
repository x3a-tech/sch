package sch

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/x3a-tech/configo"
	"time"
)

func NewClient(ctx context.Context, cfg *configo.Database) (conn driver.Conn, err error) {
	err = try(func() error {
		ctx, cancel := context.WithTimeout(ctx, cfg.AttemptDelay)
		defer cancel()

		var err error
		conn, err = clickhouse.Open(&clickhouse.Options{
			Addr: []string{fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)},
			Auth: clickhouse.Auth{
				Database: cfg.Name,
				Username: cfg.User,
				Password: cfg.Password,
			},
		})
		if err != nil {
			return fmt.Errorf("ошибка при подключении к базе данных: %w", err)
		}

		err = conn.Ping(ctx)
		if err != nil {
			return fmt.Errorf("ошибка при проверке подключения к базе данных: %w", err)
		}
		return nil
	}, cfg.MaxAttempts, cfg.AttemptDelay)

	if err != nil {
		return nil, fmt.Errorf("не удалось подключиться к базе данных после %v попыток: %w", cfg.MaxAttempts, err)
	}

	return conn, nil
}

func try(fn func() error, attempts int, delay time.Duration) (err error) {
	for attempts > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attempts--
			continue
		}
		return nil
	}
	return
}
