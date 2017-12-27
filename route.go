package main

import (
	"github.com/labstack/echo"

	"github.com/h-tko/go_oauth/handler"
)

func route(e *echo.Echo) {
	e.GET("/auth", handler.Auth)
	e.GET("/callback", handler.Callback)
}
