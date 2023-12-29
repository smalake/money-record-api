package route

import (
	"log"

	"github.com/smalake/money-record-api/internal/appmodel"
	"github.com/smalake/money-record-api/internal/service"
	"github.com/smalake/money-record-api/pkg/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetRoute(e *echo.Echo) {

	mc, err := mysql.NewClient()
	if err != nil {
		log.Fatalf("[FATAL]: %+v", err)
	}
	defer mc.Close()

	appModel := appmodel.New(mc)
	service := service.New(appModel)
	e.POST("/login-mail", service.LoginMail)

	api := e.Group("/api/v1")
	api.Use(middleware.JWT([]byte("secret")))
}
