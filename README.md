# UnClutter

Project Repository for the virtual wardrobe application UnClutter.

## To-Get-Started

### Requirements
- Golang 1.25.4
- PostgreSQL 18.0

### Install migration and code generation tool
```go install github.com/pressly/goose/v3/cmd/goose@latest```

```go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest```

### Migrate up using the migration guide below

### Initialize .env
```Copy .env.example as .env and fill in the required values```

### Run the app
```go run ./services/unc-service/```

## Migration

### Migrate up
```goose -dir migrations/unc-service/schema postgres "user=postgres password=<db_password> dbname=unc_db sslmode=disable" up```

### Migrate down
```goose -dir migrations/unc-service/schema postgres "user=postgres password=<db_password> dbname=unc_db sslmode=disable" down```

### Create migration
```goose -s -dir migrations/unc-service/schema create <migration_name> sql```
