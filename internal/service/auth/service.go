package auth

import (
	"github.com/labstack/echo/v4"

	"github.com/smalake/money-record-api/internal/appmodel"
)

type Service struct {
	appModel appmodel.AppModel
}

func New(appModel *appmodel.AppModel) *Service {
	return &Service{appModel: *appModel}
}

func (s *Service) LoginMail(ctx echo.Context) error {
	s.appModel.MysqlCli.DB.Exec("SELECT * FROM users")

	return nil
}
