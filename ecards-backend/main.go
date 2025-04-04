package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// 中间件
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 用户登录验证中间件
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 检查用户是否登录
			if !isUserLoggedIn(c) && c.Path() != "/login" {
				return c.Redirect(http.StatusFound, "/noAuth")
			}
			return next(c)
		}
	})

	// 路由
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// 登录路由
	e.GET("/login", func(c echo.Context) error {
		return c.String(http.StatusOK, "You have no permission to request!")
	})

	e.GET("/noAuth", func(c echo.Context) error {
		return c.String(http.StatusOK, "You have no permission to request!")
	})

	// 添加导入路由
	e.POST("/import", func(c echo.Context) error {
		fmt.Println("Received import request")
		var request struct {
			Text string `json:"text"`
		}
		if err := c.Bind(&request); err != nil {
			fmt.Println("Error binding request:", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}
		// 打印接收到的文本内容
		fmt.Println("Received text:", request.Text)
		return c.JSON(http.StatusOK, map[string]string{"message": "Import successful"})
	})

	// 启动服务器
	e.Logger.Fatal(e.Start(":8081"))
}

// 模拟用户登录验证
func isUserLoggedIn(c echo.Context) bool {
	// 这里可以根据实际需求实现用户登录验证逻辑
	// 例如检查session或token
	return false
}
