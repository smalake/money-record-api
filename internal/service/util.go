package service

import (
	"github.com/labstack/echo/v4"

	"github.com/smalake/money-record-api/pkg/structs"
)

func ResponseHandler(ctx echo.Context, res structs.HttpResponse) error {
	if res.Error != nil {
		return ctx.JSON(res.Code, res.Error.Error())
	}
	// 成功した場合のレスポンスを設定
	return ctx.JSON(res.Code, res.Data)
}
