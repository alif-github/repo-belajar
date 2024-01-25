package config

import (
	"database/sql"
	_ "github.com/jackc/pgx/stdlib"
	"log"
	"sync"
)

var (
	ServerAttribute serverAttribute
	once            sync.Once
)

type serverAttribute struct {
	DBConnection *sql.DB
}

func SetServerAttribute() {
	var (
		err      error
		instance *sql.DB
	)

	once.Do(func() {
		instance, err = sql.Open("pgx", "user=postgres password=paramadaksa dbname=grpcTEST sslmode=disable host=localhost port=5432 search_path='grpc'")
		if err != nil {
			log.Fatalln("Failed open connection database -> ", err.Error())
		}

		ServerAttribute.DBConnection = instance
	})
}
