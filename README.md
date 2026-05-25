# UnClutter

Project Repository for the virtual wardrobe application UnClutter.

## To-Get-Started

### Requirements
- Golang 1.26
- PostgreSQL 18.0

### Install migration, code generation, hot-reload, and command tools
```go install github.com/pressly/goose/v3/cmd/goose@latest```

```go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest```

```go install github.com/air-verse/air@latest```

```go install github.com/go-task/task/v3/cmd/task@latest```

### Initialize .env
```Copy .env.example as .env and fill in the required values```

### Changing environment
```Change dotenv value in Taskfile.yml to .env.dev for dev environment as example```

### Run the app
```task dev```

### Build the app and run
```task build & task run```

### Test the app
```task test```

## Manual migration

### Migrate create
```task migrate:create -- <migration_name>```

### Migrate up
```task migrate:up```

### Migrate down
```task migrate:down -- <migration_number>```

### Migrate up-to
```task migrate:up-to -- <migration_number>```

### Migrate down-to
```task migrate:down-to```
