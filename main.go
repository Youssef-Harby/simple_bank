package main

import (
	"database/sql"
	"log"

	"github.com/a-saber-abdelmosen/simplebank/api"
	db "github.com/a-saber-abdelmosen/simplebank/db/sqlc"
	"github.com/a-saber-abdelmosen/simplebank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("connot load app configuration", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("connot connect to the database", err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Connot start the server", err)
	}
}
