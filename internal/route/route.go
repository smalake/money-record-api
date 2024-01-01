package route

import (
	"github.com/smalake/money-record-api/internal/appmodel"
	"github.com/smalake/money-record-api/internal/service"
	"github.com/smalake/money-record-api/pkg/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetRoute(e *echo.Echo) {

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339_nano}] method=${method}, uri=${uri}, status=${status}\n",
	}))

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

	api := e.Group("/api/v1")
	api.Use(middleware.JWT([]byte("secret")))
	api.POST("/logout", service.LogoutHandler)
}
