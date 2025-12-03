package setup

import (
	"context"
	"fmt"
	"log"
	"unc/services/unc-service/config"
	"unc/services/unc-service/domain/repository"

	"github.com/jackc/pgx/v5"
)

func setupDB(ctx context.Context) repository.DBTX {
	conn, err := pgx.Connect(ctx, fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		config.GetConfig().DatabaseUser,
		config.GetConfig().DatabasePassword,
		config.GetConfig().DatabaseHost,
		config.GetConfig().DatabasePort,
		config.GetConfig().DatabaseName,
	))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err = conn.Close(ctx)
		if err != nil {
			log.Fatalf("Unable to close connection: %v", err)
		}
	}(conn, ctx)

	return conn
}

func setupRepository(conn repository.DBTX) repository.Querier {
	return repository.New(conn)
}
