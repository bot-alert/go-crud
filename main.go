package main

import (
	"GO/database"
	"GO/server"
	"log"
)

func main() {
	database.StartDatabase()
	log.Fatal(server.StartServer())
}
