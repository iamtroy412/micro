package main

import (
	"log"

	"github.com/iamtroy412/micro/internal/database"
	"github.com/iamtroy412/micro/internal/server"
)

func main() {
	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("failed to initialize Database Client: %s", err)
	}
	srv := server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
