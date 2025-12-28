package main

import (
	"flexy/delivery/httpserver"
	"flexy/repository/migrator"
	"flexy/repository/sqlite"
)

func main() {

	config := sqlite.Config{
		FilePath: "./database.db",
	}
	migrator := migrator.New(config)
	migrator.Up()

	server := httpserver.New()
	server.Serve()
}
