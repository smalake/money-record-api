package service

import "github.com/labstack/echo/v4"

type ServerInterface interface {

	// (POST /auth-code)
	AuthCode(ctx echo.Context) error

	// (GET /health-check)
	HealthCheck(ctx echo.Context) error

	// (GET /login-check)
	LoginCheck(ctx echo.Context) error
	// Googleアカウントでログイン
	// (POST /login-google)
	LoginGoogle(ctx echo.Context) error
	// メールアドレスでログイン
	// (POST /login-mail)
	LoginMail(ctx echo.Context) error
	// ログアウト処理
	// (GET /logout)
	Logout(ctx echo.Context) error

	// (POST /register)
	RegisterUser(ctx echo.Context) error

	// (GET /resend-code)
	ResendCode(ctx echo.Context) error

	// (POST /send-mail)
	SendMail(ctx echo.Context) error
}
