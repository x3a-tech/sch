# SPG - Storage PG

Данный пакет служит для упрощения подключения к клиенту pgx, выполнения миграций и чтения конфига базы данных

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

Путь `md/migrate/run/main.go`

```go
package main

import (
	"github.com/x3a-tech/spg"
	"path/internal/config"
)

func main() {
	cfg := config.MustLoad()
	spg.MigrateRun(cfg.Database)
}

```

### Создание миграций
Путь `cmd/migrate/create/main.go`

```go
package main

import (
	"github.com/x3a-tech/spg"
	"path/internal/config"
)

func main() {
	cfg := config.MustLoad()
	spg.MigrateCreate(cfg.Database)
}

```
