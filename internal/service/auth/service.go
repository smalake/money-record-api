package auth

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/smalake/money-record-api/internal/appmodel"
	"github.com/smalake/money-record-api/pkg/mysql"
	"github.com/smalake/money-record-api/pkg/structs"
	"github.com/smalake/money-record-api/pkg/user"
)

type Service struct {
	appModel appmodel.AppModel
}

func New(appModel *appmodel.AppModel) *Service {
	return &Service{appModel: *appModel}
}

// メールアドレスでログイン
func (s *Service) LoginMail(ctx echo.Context) structs.HttpResponse {
	// POSTからログイン情報を取得
	u := new(user.LoginMailRequest)
	if err := ctx.Bind(u); err != nil {
		return structs.HttpResponse{Code: 400, Error: err}
	}

	query := mysql.LoginMail
	var uid user.LoginMailInfo
	err := s.appModel.MysqlCli.DB.Get(&uid, query, u.Email)
	if err != nil {
		return structs.HttpResponse{Code: 500, Error: err}
	}

	if err := compareHashAndPassword(uid.Password, u.Password); err != nil {
		// 認証エラー
		return structs.HttpResponse{Code: 401, Error: err}
	}

	return structs.HttpResponse{Code: 200, Data: "token"}
}

// ハッシュ化されたパスワードとの比較(returnがnilならログイン成功)
func compareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
