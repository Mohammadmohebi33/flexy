package main

import (
	"flexy/repository/migrator"
	"flexy/repository/sqlite"
)

func main() {

	config := sqlite.Config{
		FilePath: "./database.db",
	}
	migrator := migrator.New(config)
	migrator.Up()
}
