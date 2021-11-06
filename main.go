package main

import (
	"books-api/database"
	"books-api/database/migrations"
	"books-api/server"
)

func main() {
	database.StartDB()
	migrations.RunMigrations(database.GetDatabase())

	server := server.NewServer()
	server.Run()
}
