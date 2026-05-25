package setup

import (
	"database/sql"
	"fmt"
	"log"
	"unc/services/unc-service/config"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func migrateGoose() {
	db, err := sql.Open("pgx", fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		config.GetConfig().DatabaseUser,
		config.GetConfig().DatabasePassword,
		config.GetConfig().DatabaseHost,
		config.GetConfig().DatabasePort,
		config.GetConfig().DatabaseName,
	))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}

	if err = goose.Up(db, "migrations/unc-service/schema"); err != nil {
		log.Fatal(err)
	}
}
