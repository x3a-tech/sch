# SCH - Storage CH

Данный пакет служит для упрощения подключения к клиенту clickhouse, выполнения миграций и чтения конфига базы данных

Предварительно необходимо иметь конфиг файл и конфиг структуру "./internal/config/config.go" configo.Config

## Таски для файла Taskfile.yml

```yaml
version: '3'

tasks:
  migrate:
    cmds:
      - task: migrate:{{.CLI_ARGS}}

  migrate:run:
    desc: Запуск миграций
    cmds:
      - go run ./cmd/migrate/run/main.go -config=./config/local.yml

  migrate:create:
    desc: Создание миграций
    cmds:
      - go run ./cmd/migrate/create/main.go -config=./config/local.yml {{.NAME}}
    vars:
      NAME:
        sh: |
          echo {{.CLI_ARGS}}
```

## cmd

Предварительно необходимо создать файлы и при копировании `"path/internal/config"` `path` заменить на название своего модуля

### Запуск миграций

Путь `cmd/migrate/run/main.go`

```go
package main

import (
	"github.com/x3a-tech/sch"
	"path/internal/config"
)

func main() {
	cfg := config.MustLoad()
	sch.MigrateRun(cfg.Database)
}

```

### Создание миграций
Путь `cmd/migrate/create/main.go`

```go
package main

import (
	"github.com/x3a-tech/sch"
	"path/internal/config"
)

func main() {
	cfg := config.MustLoad()
	sch.MigrateCreate(cfg.Database)
}

```
