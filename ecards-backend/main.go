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
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"}, // 允许所有域名访问，生产环境应限制为特定域名
	// 	AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	// }))

	// 路由
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
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
