package http

import (
	"github.com/PongponZ/go-book-store/internal/delivery/http/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handlers struct {
	UserHandler handler.IUserHandler
}

func RegisterRoutes(e *echo.Echo, h *Handlers) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/users", h.UserHandler.Create)
	e.POST("/users", h.UserHandler.Create)
	e.GET("/users/:userID", h.UserHandler.FindByID)
	e.PUT("/users/:userID", h.UserHandler.Update)
	e.DELETE("/users/:userID", h.UserHandler.DeleteByID)
}
