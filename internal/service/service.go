package service

import (
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
func (s Service) LoginGoogleHandler(ctx echo.Context) error {
	service := auth.New(&s.appModel)
	res := service.LoginGoogle(ctx)

	return ResponseHandler(ctx, res)
}

func (s Service) LoginMailHandler(ctx echo.Context) error {
	service := auth.New(&s.appModel)
	res := service.LoginMail(ctx)

	return ResponseHandler(ctx, res)
}

func (s Service) LogoutHandler(ctx echo.Context) error {
	service := auth.New(&s.appModel)
	res := service.Logout(ctx)

	return ResponseHandler(ctx, res)
}

func (s Service) RegisterUserHandler(ctx echo.Context) error {
	service := auth.New(&s.appModel)
	res := service.RegisterUser(ctx)

	return ResponseHandler(ctx, res)
}

func (s Service) LoginCheckHandler(ctx echo.Context) error {

	return nil
}

func (s Service) AuthCodeHandler(ctx echo.Context) error {

	return nil
}

func (s Service) ResendCodeHandler(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

func (s Service) SendMailHandler(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

// HealthCheck
func (s Service) HealthCheckHandler(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}
