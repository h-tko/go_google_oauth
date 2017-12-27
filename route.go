package main

import (
	"github.com/labstack/echo"

	"github.com/mm-lvgs/mss_auth/handler"
)

func route(e *echo.Echo) {
	e.GET("/auth", handler.Auth)
	e.GET("/callback", handler.Callback)
}
