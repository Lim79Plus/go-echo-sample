package api

import (
	"github.com/Lim79Plus/go_echo_sample/api/handler"
	"github.com/labstack/echo"
)

// MainGroup set main group router for handler
func MainGroup(e *echo.Echo) {
	e.GET("/hello", handler.MainPage())
}
