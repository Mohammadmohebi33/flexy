package main

import (
	"flexy/config"
	"flexy/delivery/httpserver"
	"flexy/repository/sqlite"
	"flexy/repository/sqlite/usersqlite"
	"flexy/service/authservice"
	"flexy/service/userservice"
)

func main() {

	dbConfig := sqlite.Config{
		FilePath: "./database.db",
	}
	//migrator := migrator.New(config)
	//migrator.Up()
	//

	jwtConfig := authservice.Config{
		SignKey:               config.JwtSignKey,
		AccessExpirationTime:  config.AccessTokenExpireDuration,
		RefreshExpirationTime: config.RefreshTokenExpireDuration,
		AccessSubject:         config.AccessTokenSubject,
		RefreshSubject:        config.RefreshTokenSubject,
	}

	database := sqlite.New(dbConfig)
	userRepo := usersqlite.New(database)
	authService := authservice.New(jwtConfig)
	userService := userservice.New(authService, userRepo)
	server := httpserver.New(userService)
	server.Serve()
}
