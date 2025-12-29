package httpserver

import (
	"flexy/delivery/httpserver/authhandler"
	"flexy/service/authservice"
	"fmt"

	"github.com/labstack/echo/v4"
)

type Server struct {
	authHandler authhandler.Handler
	Router      *echo.Echo
}

func New(authSvc authservice.Service) Server {
	return Server{
		Router:      echo.New(),
		authHandler: authhandler.New(authSvc),
	}
}

func (s Server) Serve() {
	// Routes
	s.Router.GET("/health-check", s.healthCheck)
	s.authHandler.SetRoutes(s.Router)

	// Start server
	address := fmt.Sprintf(":%d", 8080)
	fmt.Printf("start echo server on %s\n", address)
	if err := s.Router.Start(address); err != nil {
		fmt.Println("router start error", err)
	}
}
