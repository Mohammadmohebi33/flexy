package authhandler

import "github.com/labstack/echo/v4"

func (h Handler) SetRoutes(e *echo.Echo) {
	userGroup := e.Group("/auth")

	userGroup.POST("/register", h.userRegister)
}
