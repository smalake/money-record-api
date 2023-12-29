package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/smalake/money-record-api/internal/appmodel"
	"github.com/smalake/money-record-api/internal/service/auth"
)

type Service struct {
	appModel appmodel.AppModel
}

func New(am *appmodel.AppModel) *Service {
	return &Service{appModel: *am}
}

// Auth関連
func (s Service) LoginGoogle(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

func (s Service) LoginMail(ctx echo.Context) error {
	service := auth.New(&s.appModel)

	return ctx.JSON(http.StatusOK, service.LoginMail)
}

func (s Service) Logout(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

func (s Service) RegisterUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

func (s Service) LoginCheck(ctx echo.Context) error {

	return nil
}

func (s Service) AuthCode(ctx echo.Context) error {

	return nil
}

func (s Service) ResendCode(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

func (s Service) SendMail(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

// HealthCheck
func (s Service) HealthCheck(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}
