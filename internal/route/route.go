package route

import (
	"errors"
	"fmt"
	"strings"

	"github.com/smalake/money-record-api/internal/appmodel"
	"github.com/smalake/money-record-api/internal/service"
	"github.com/smalake/money-record-api/pkg/mysql"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetRoute(e *echo.Echo) {

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339_nano}] method=${method}, uri=${uri}, status=${status}, ${error}\n",
	}))
	e.Use(middleware.Recover())

	mc, err := mysql.NewClient()
	if err != nil {
		e.Logger.Fatalf("[FATAL]: %+v", err)
	}
	// defer mc.Close()

	appModel := appmodel.New(mc)
	service := service.New(appModel)
	e.POST("/login-mail", service.LoginMailHandler)
	e.POST("/login-google", service.LoginGoogleHandler)
	e.POST("/register", service.RegisterUserHandler)
	e.POST("/auth-code", service.AuthCodeHandler)

	api := e.Group("/api/v1")
	api.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("secret"),
	}))

	// JWT認証
	api.Use(JWTMiddleware)
	api.POST("/logout", service.LogoutHandler)
}

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return errors.New("invalid or missing JWT")
		}
		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("secret"), nil
		})
		if err != nil {
			return err
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// JWTからユーザーIDを取得
			uid, ok := claims["id"]
			if !ok {
				return errors.New("failed to get uid from JWT")
			}
			// ユーザーIDをContextに設定
			c.Set("uid", uid)
			return next(c)
		} else {
			return errors.New("invalid or expired JWT")
		}
	}
}
