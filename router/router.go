package router

import (
	"github.com/Lim79Plus/go-echo-sample/api/middleware"
	"github.com/Lim79Plus/go_echo_sample/api"
	"github.com/labstack/echo"
)

// New create instance of Echo and set router
func New() *echo.Echo {
	// Echoのインスタンス作る
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア
	middleware.SetMainMiddleware(e)

	// ルーティング
	api.MainGroup(e)
	// e.GET("/hello", mainPage())

	return e
}
