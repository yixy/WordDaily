package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var jwtSecret = []byte("your-secret-key") // 从环境变量中读取更安全

func generateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 1).Unix(), // Token有效期为24小时
	})
	return token.SignedString(jwtSecret)
}

func main() {
	e := echo.New()

	// 中间件
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 用户登录验证中间件
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 检查用户是否登录
			if isUserLoggedIn(c) || c.Path() == "/api/login" {
				return next(c)
			}
			return c.JSON(http.StatusUnauthorized, map[string]bool{"success": false})
			//return c.Redirect(http.StatusFound, "/api/auth")
		}
	})

	api:= e.Group("/api")

	// 路由
	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// 登录路由
	api.POST("/login", func(c echo.Context) error {
		var request struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}

		// 模拟登录验证逻辑
		if request.Username == "admin" && request.Password == "password" {
			// 登录成功，设置用户登录状态
			token, err := generateToken(request.Username)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
			}
			return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "authToken": token})
		} else {
			return c.JSON(http.StatusUnauthorized, map[string]bool{"success": false})
		}
	})

	api.GET("/noAuth", func(c echo.Context) error {
		return c.String(http.StatusOK, "You have no permission to request!")
	})

	// 添加导入路由
	api.POST("/import", func(c echo.Context) error {
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
	tokenString := c.Request().Header.Get("Authorization")
	if tokenString == "" {
		return false
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return false
	}
	return true
}
