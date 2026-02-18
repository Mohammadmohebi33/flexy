package main

import (
	"flexy/delivery/httpserver"
	"flexy/repository/sqlite"
	"flexy/repository/sqlite/usersqlite"
	"flexy/service/userservice"
)

func main() {

	config := sqlite.Config{
		FilePath: "./database.db",
	}
	//migrator := migrator.New(config)
	//migrator.Up()
	//

	database := sqlite.New(config)
	authRepo := usersqlite.New(database)
	authService := userservice.New(authRepo)
	server := httpserver.New(authService)
	server.Serve()
}
