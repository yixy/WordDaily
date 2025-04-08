package main

import (
	"fmt"
	"net/http"
	"time"
	"worddaily-backend/logger"
	"worddaily-backend/model"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	jwtSecret = []byte("your-secret-key") // 从环境变量中读取更安全
	store     = sessions.NewCookieStore([]byte("your-secret-key"))
)

func init() {
	// 初始化 Viper 配置
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		logger.LogError("Failed to read config file:", zap.Error(err))
		panic(err)
	}

	// 初始化数据库
	dsn := viper.GetString("database.dsn")
	if err := model.InitDB(dsn); err != nil {
		logger.LogError("Failed to initialize database:", zap.Error(err))
		panic(err)
	}
}

func generateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token有效期
	})
	return token.SignedString(jwtSecret)
}

func main() {
	e := echo.New()

	// 中间件
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 自定义日志记录中间件
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 记录请求日志
			logger.LogInfo("Request Log:",
				zap.String("method", c.Request().Method),
				zap.String("path", c.Request().URL.Path),
				zap.Any("headers", c.Request().Header),
				zap.Any("body", c.Request().Body),
			)

			// 继续处理请求
			err := next(c)

			// 记录响应日志
			if err != nil {
				logger.LogError("Response Log:",
					zap.Int("status", c.Response().Status),
					zap.Any("headers", c.Response().Header()),
					zap.Error(err),
				)
			} else {
				logger.LogInfo("Response Log:",
					zap.Int("status", c.Response().Status),
					zap.Any("headers", c.Response().Header()),
				)
			}

			return err
		}
	})

	// 用户登录验证中间件
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 检查用户是否登录
			if isUserLoggedIn(c) || c.Path() == "/api/login" {
				return next(c)
			}
			return c.JSON(http.StatusUnauthorized, map[string]bool{"success": false})
		}
	})

	api := e.Group("/api")

	// 路由
	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// 登出
	api.POST("/logout", func(c echo.Context) error {
		session, err := store.Get(c.Request(), "sessionID")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]bool{"success": false})
		}

		// 作废 session 信息
		session.Options.MaxAge = -1
		if err := session.Save(c.Request(), c.Response()); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to save session"})
		}

		return c.JSON(http.StatusUnauthorized, map[string]bool{"success": true})
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

		// 查询数据库验证用户名和密码
		user, err := model.GetUserByUsername(request.Username)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]bool{"success": false})
		}
		if user.UserPwd != request.Password {
			return c.JSON(http.StatusUnauthorized, map[string]bool{"success": false})
		}

		// 登录成功，设置用户登录状态
		token, err := generateToken(request.Username)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
		}

		// 设置 session 信息
		session, err := store.Get(c.Request(), "sessionID")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get session"})
		}
		session.Values["username"] = request.Username
		session.Values["sessionID"] = generateSessionID()
		session.Options.MaxAge = 86400 // 24小时
		if err := session.Save(c.Request(), c.Response()); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to save session"})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "authToken": token})
	})

	// 添加导入路由
	api.POST("/import", func(c echo.Context) error {
		var request struct {
			Text string `json:"text"`
		}
		if err := c.Bind(&request); err != nil {
			logger.LogError("Error binding request:", zap.Error(err))
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}
		// 打印接收到的文本内容
		fmt.Println("Received text:", request.Text)
		return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "message": "Import successful"})
	})

	// 新增用户修改接口
	api.POST("/usermod", func(c echo.Context) error {
		var request struct {
			Username string `json:"username"`
			Headshot string `json:"headshot"`
		}
		if err := c.Bind(&request); err != nil {
			logger.LogError("Error binding request:", zap.Error(err))
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}

		// 调用 user.go 中的方法更新用户头像
		err := model.UpdateUserHeadshot(request.Headshot, request.Username)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "message": "Headshot updated successfully"})
	})

	api.POST("/userquery", func(c echo.Context) error {
		var request struct {
			Username string `json:"username"`
		}
		if err := c.Bind(&request); err != nil {
			logger.LogError("Error binding request:", zap.Error(err))
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}

		// 调用 user.go 中的方法
		user, err := model.GetUserByUsername(request.Username)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
		}

		// 返回用户信息
		return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "user": user})
	})

	// 启动服务器
	e.Logger.Fatal(e.Start(":8081"))
}

// 模拟用户登录验证
func isUserLoggedIn(c echo.Context) bool {
	tokenString := c.Request().Header.Get("Authorization")
	logger.LogDebug("TokenString:", zap.String("tokenString", tokenString))

	// 检查是否包含 "Bearer " 前缀
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:] // 去掉 "Bearer " 前缀
	} else {
		return false // 如果没有 "Bearer " 前缀，直接返回 false
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return false
	}
	return true
}

// 辅助函数：生成 sessionID
func generateSessionID() string {
	return uuid.New().String()
}
