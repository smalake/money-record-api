package auth

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"net/smtp"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/smalake/money-record-api/internal/appmodel"
	"github.com/smalake/money-record-api/internal/env"
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

	// トークンを発行
	token, err := issueToken(uid.ID)
	if err != nil {
		return structs.HttpResponse{Code: 500, Error: err}
	}

	return structs.HttpResponse{Code: 200, Data: map[string]string{"accessToken": token}}
}

// Googleアカウントでログイン
func (s *Service) LoginGoogle(ctx echo.Context) structs.HttpResponse {
	// POSTからログイン情報を取得
	u := new(user.LoginGoogleRequest)
	if err := ctx.Bind(u); err != nil {
		return structs.HttpResponse{Code: 400, Error: err}
	}

	query := mysql.LoginGoogle
	var uid user.UserID
	err := s.appModel.MysqlCli.DB.Get(&uid, query, u.Email)
	if err != nil {
		return structs.HttpResponse{Code: 500, Error: err}
	}

	// トークンを発行
	token, err := issueToken(uid.ID)
	if err != nil {
		return structs.HttpResponse{Code: 500, Error: err}
	}

	return structs.HttpResponse{Code: 200, Data: map[string]string{"accessToken": token}}
}

// ユーザ登録
func (s *Service) RegisterUser(ctx echo.Context) structs.HttpResponse {
	// POSTからユーザ情報を取得
	u := new(user.RegisterUserRequest)
	if err := ctx.Bind(u); err != nil {
		return structs.HttpResponse{Code: 400, Error: err}
	}

	// パスワードをハッシュ化
	password, err := passwordEncrypt(u.Password)
	if err != nil {
		return structs.HttpResponse{Code: 500, Error: err}
	}

	// 認証コードを発行
	newSeed := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := newSeed.Intn(1000000) + 100000

	// メール送信の設定を取得
	mailConfig, err := env.GetMailConfig()
	if err != nil {
		ctx.Logger().Errorf("[FATAL] %v", err)
		return structs.HttpResponse{Code: 500, Error: err}
	}
	auth := smtp.PlainAuth("", mailConfig.Username, mailConfig.Password, mailConfig.Host)
	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: 【Money Record】認証コード\r\n\r\n認証コード: %d", u.Email, code))

	// メールアドレスが既に登録されていないかチェック
	query := mysql.CheckEmail
	var authCode int
	err = s.appModel.MysqlCli.DB.Get(&authCode, query, u.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			// メールアドレスが未登録の場合は、新規登録
			tx, err := s.appModel.MysqlCli.DB.Beginx()
			if err != nil {
				return structs.HttpResponse{Code: 500, Error: err}
			}
			// ユーザ作成
			query = mysql.CreateUser
			res, err := tx.Exec(query, u.Email, password, u.Name, u.RegisterType, code)
			if err != nil {
				tx.Rollback()
				return structs.HttpResponse{Code: 500, Error: err}
			}
			// ユーザIDを取得
			uid, err := res.LastInsertId()
			if err != nil {
				tx.Rollback()
				return structs.HttpResponse{Code: 500, Error: err}
			}
			// グループ作成
			query = mysql.CreateGroup
			res, err = tx.Exec(query, uid)
			if err != nil {
				tx.Rollback()
				return structs.HttpResponse{Code: 500, Error: err}
			}
			// グループIDを取得
			gid, err := res.LastInsertId()
			if err != nil {
				tx.Rollback()
				return structs.HttpResponse{Code: 500, Error: err}
			}
			// グループIDをユーザに紐付ける
			query = mysql.UpdateGroup
			_, err = tx.Exec(query, gid, uid)
			if err != nil {
				tx.Rollback()
				return structs.HttpResponse{Code: 500, Error: err}
			}
			// メールアドレスによる登録の場合はメール送信
			if u.RegisterType == 1 {
				// メール送信処理
				if err := smtp.SendMail(fmt.Sprintf("%s:%d", mailConfig.Host, mailConfig.Port), auth, mailConfig.From, []string{u.Email}, msg); err != nil {
					ctx.Logger().Errorf("[FATAL] %v", err)
					return structs.HttpResponse{Code: 500, Error: err}
				}
			}

			// メールアドレスによる登録でない場合は認証コードを更新
			query = mysql.UpdateAuthCode
			_, err = tx.Exec(query, 0, u.Email)
			if err != nil {
				tx.Rollback()
				return structs.HttpResponse{Code: 500, Error: err}
			}

			err = tx.Commit()
			if err != nil {
				return structs.HttpResponse{Code: 500, Error: err}
			}
			return structs.HttpResponse{Code: 200}
		} else {
			// SELECT文の実行に失敗した場合
			ctx.Logger().Errorf("[FATAL] %v", err)
			return structs.HttpResponse{Code: 500, Error: err}
		}
	}
	if authCode == 0 {
		// すでに登録されているメールアドレスが使用された場合409エラー
		ctx.Logger().Errorf("[FATAL] %v", err)
		return structs.HttpResponse{Code: 409, Error: errors.New("already registered email")}
	} else {
		// すでに登録されているメールアドレスだが未認証の場合は認証コードを再送信
		query := mysql.ResendAuthCode
		_, err = s.appModel.MysqlCli.DB.Exec(query, code, password, u.Name, u.Email)
		if err != nil {
			ctx.Logger().Errorf("[FATAL] %v", err)
			return structs.HttpResponse{Code: 500, Error: err}
		}
		// メール送信
		if err := smtp.SendMail(fmt.Sprintf("%s:%d", mailConfig.Host, mailConfig.Port), auth, mailConfig.From, []string{u.Email}, msg); err != nil {
			ctx.Logger().Errorf("[FATAL] %v", err)
			return structs.HttpResponse{Code: 500, Error: err}
		}
		return structs.HttpResponse{Code: 200}
	}
}

// ログアウト
func (s *Service) Logout(ctx echo.Context) structs.HttpResponse {
	return structs.HttpResponse{Code: 200}
}

// ログインチェック(親か子かを判定も行う)
func (s *Service) LoginCheck(ctx echo.Context) structs.HttpResponse {
	uid := ctx.Get("uid")
	// ユーザIDから親か子かを判定
	var count int
	query := mysql.LoginCheck
	err := s.appModel.MysqlCli.DB.Get(&count, query, uid)
	if err != nil {
		return structs.HttpResponse{Code: 500, Error: err}
	}
	if count == 0 {
		// 子の場合
		return structs.HttpResponse{Code: 200, Data: map[string]bool{"parent": false}}
	} else {
		// 親の場合
		return structs.HttpResponse{Code: 200, Data: map[string]bool{"parent": true}}
	}
}

// トークン発行
func issueToken(id int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["iat"] = time.Now().Unix()                      // 発行時間
	claims["exp"] = time.Now().Add(time.Hour * 720).Unix() // 有効期限(30日)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}

// パスワードのハッシュ化
func passwordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// ハッシュ化されたパスワードとの比較(returnがnilならログイン成功)
func compareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
