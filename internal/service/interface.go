package service

import "github.com/labstack/echo/v4"

type ServerInterface interface {

	// (POST /auth-code)
	AuthCodeHandler(ctx echo.Context) error

	// (GET /health-check)
	HealthCheckHandler(ctx echo.Context) error

	// (GET /login-check)
	LoginCheckHandler(ctx echo.Context) error
	// Googleアカウントでログイン
	// (POST /login-google)
	LoginGoogleHandler(ctx echo.Context) error
	// メールアドレスでログイン
	// (POST /login-mail)
	LoginMailHandler(ctx echo.Context) error
	// ログアウト処理
	// (GET /logout)
	LogoutHandler(ctx echo.Context) error

	// (POST /register)
	RegisterUserHandler(ctx echo.Context) error

	// (GET /resend-code)
	ResendCodeHandler(ctx echo.Context) error

	// (POST /send-mail)
	SendMailHandler(ctx echo.Context) error

	// (GET /memo)
	GetMemoAllHandler(ctx echo.Context) error

	// (GET /memo/:id)
	GetMemoOneHandler(ctx echo.Context) error

	// (POST /memo)
	CreateMemoHandler(ctx echo.Context) error

	// (PUT /memo)
	UpdateMemoHandler(ctx echo.Context) error

	// (DELETE /memo)
	DeleteMemoHandler(ctx echo.Context) error
}
