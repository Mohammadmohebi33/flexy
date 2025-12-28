package httpserver

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Router *echo.Echo
}

func New() Server {
	return Server{
		Router: echo.New(),
	}
}

func (s Server) Serve() {
	// Routes
	s.Router.GET("/health-check", s.healthCheck)

	// Start server
	address := fmt.Sprintf(":%d", 8080)
	fmt.Printf("start echo server on %s\n", address)
	if err := s.Router.Start(address); err != nil {
		fmt.Println("router start error", err)
	}
}
