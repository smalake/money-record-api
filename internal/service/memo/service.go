package memo

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/smalake/money-record-api/internal/appmodel"
	memos "github.com/smalake/money-record-api/pkg/memo"
	"github.com/smalake/money-record-api/pkg/mysql"
	"github.com/smalake/money-record-api/pkg/structs"
)

type Service struct {
	appModel appmodel.AppModel
}

func New(am *appmodel.AppModel) *Service {
	return &Service{appModel: *am}
}

func (s *Service) GetMemoAll(ctx echo.Context) structs.HttpResponse {
	r := new(memos.CreateMemoRequest)
	if err := ctx.Bind(r); err != nil {
		ctx.Logger().Errorf("[FATAL] failed to bind: %+v", err)
		return structs.HttpResponse{Code: 400, Error: err}
	}
	uid := ctx.Get("uid")

	query := mysql.GetMemoAll
	var memo memos.GetMemoResponse
	err := s.appModel.MysqlCli.DB.Select(&memo, query, uid)
	if err != nil {
		ctx.Logger().Errorf("[FATAL] failed to get all memos: %+v", err)
		return structs.HttpResponse{Code: 500, Error: err}
	}
	return structs.HttpResponse{Code: 200, Data: memo}
}

func (s *Service) GetMemoOne(ctx echo.Context) structs.HttpResponse {
	uid := ctx.Get("uid")
	id := ctx.Param("id")
	if id == "" {
		ctx.Logger().Errorf("[ERROR] memo id is empty")
		return structs.HttpResponse{Code: 400, Error: fmt.Errorf("[ERROR] memo id is empty")}
	}

	query := mysql.GetMemoOne
	var memo memos.OneMemo
	err := s.appModel.MysqlCli.DB.Get(&memo, query, id, uid)
	if err != nil {
		ctx.Logger().Errorf("[FATAL] failed to get one memo: %+v", err)
		return structs.HttpResponse{Code: 500, Error: err}
	}
	return structs.HttpResponse{Code: 200, Data: memo}
}

func (s *Service) CreateMemo(ctx echo.Context) structs.HttpResponse {
	r := new(memos.OneMemo)
	if err := ctx.Bind(r); err != nil {
		ctx.Logger().Errorf("[FATAL] failed to bind: %+v", err)
		return structs.HttpResponse{Code: 400, Error: err}
	}
	uid := ctx.Get("uid")
	query := mysql.CreateMemo
	_, err := s.appModel.MysqlCli.DB.Exec(query, uid, r.Amount, r.Partner, r.Memo, r.Date, r.Period, r.Type)
	if err != nil {
		ctx.Logger().Errorf("[FATAL] failed to create new memo: %+v", err)
		return structs.HttpResponse{Code: 500, Error: err, Message: err.Error()}
	}
	return structs.HttpResponse{Code: 200}
}

func (s *Service) UpdateMemo(ctx echo.Context) structs.HttpResponse {
	r := new(memos.OneMemo)
	if err := ctx.Bind(r); err != nil {
		ctx.Logger().Errorf("[FATAL] failed to bind: %+v", err)
		return structs.HttpResponse{Code: 400, Error: err}
	}
	uid := ctx.Get("uid")
	query := mysql.UpdateMemo
	_, err := s.appModel.MysqlCli.DB.Exec(query, r.Amount, r.Partner, r.Memo, r.Date, r.Period, r.Type, r.ID, uid)
	if err != nil {
		ctx.Logger().Errorf("[FATAL] failed to update memo: %+v", err)
		return structs.HttpResponse{Code: 500, Error: err}
	}
	return structs.HttpResponse{Code: 200}
}

func (s *Service) DeleteMemo(ctx echo.Context) structs.HttpResponse {
	r := new(memos.OneMemo)
	if err := ctx.Bind(r); err != nil {
		ctx.Logger().Errorf("[FATAL] failed to bind: %+v", err)
		return structs.HttpResponse{Code: 400, Error: err}
	}
	uid := ctx.Get("uid")
	query := mysql.DeleteMemo
	_, err := s.appModel.MysqlCli.DB.Exec(query, r.ID, uid)
	if err != nil {
		ctx.Logger().Errorf("[FATAL] failed to delete memo %+v", err)
		return structs.HttpResponse{Code: 500, Error: err}
	}
	return structs.HttpResponse{Code: 200}
}
