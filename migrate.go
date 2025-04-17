package sch

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/clickhouse"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/x3a-tech/configo"
	"os"
	"path/filepath"
	"time"
)

func MigrateRun(cfg configo.Database) {
	runMigrations(&cfg)
}

func MigrateCreate(cfg configo.Database) {
	args := make([]string, 0, 2)
	args = append(args, os.Args[2:]...)
	name := args[0]

	createMigrationFiles(cfg.MigrationPath, name)
}

func Dsn(cfg *configo.Database) string {
	return fmt.Sprintf("clickhouse://%s:%d/?username=%s&password=%s&database=%s&x-multi-statement=true", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
}

func runMigrations(cfg *configo.Database) {
	m, err := migrate.New(
		"file://"+cfg.MigrationPath,
		Dsn(cfg),
	)

	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("Нет ничего для миграции")
			return
		}
		panic(err.Error())
	} else {
		fmt.Println("Миграция успешно выполнена")
	}
}

func createMigrationFiles(dir, name string) {
	// Создание директории для миграций, если она не существует
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}

	// Получение текущего времени для префикса
	now := time.Now().Format("200601021504")

	// Имена файлов миграций
	upFileName := fmt.Sprintf("%s_%s.up.sql", now, name)
	downFileName := fmt.Sprintf("%s_%s.down.sql", now, name)

	// Пути к файлам миграций
	upFilePath := filepath.Join(dir, upFileName)
	downFilePath := filepath.Join(dir, downFileName)

	// Создание файла для применения миграции
	upFile, err := os.Create(upFilePath)
	if err != nil {
		panic(err)
	}
	upFile.Close()

	// Создание файла для отката миграции
	downFile, err := os.Create(downFilePath)
	if err != nil {
		panic(err)
	}
	downFile.Close()

	fmt.Printf("Миграция создана:\n%s\n%s\n", upFilePath, downFilePath)
}
