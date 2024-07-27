package main

import (
	"go-hipo/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echoインスタンスの作成
	e := echo.New()

	// ルートハンドラの設定
	routes.InitRoutes(e)

	// サーバの開始
	e.Start(":1323")
}
