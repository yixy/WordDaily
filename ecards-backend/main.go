package main

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "net/http"
)

func main() {
    // 创建 Echo 实例
    e := echo.New()

    // 中间件
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // 路由
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })

    // 启动服务器
    e.Logger.Fatal(e.Start(":8080"))
}