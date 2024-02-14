package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/techschool/simplebank/api"
	db "github.com/techschool/simplebank/db/sqlc"
)

const (
	dbDriver       = "postgres"
	dbSource       = "postgresql://root:12345@localhost:5432/simple_bank?sslmode=disable"
	serverAddresss = "0.0.0.0:8080"
)

func main() {
	var err error
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connected to the database", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddresss)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
