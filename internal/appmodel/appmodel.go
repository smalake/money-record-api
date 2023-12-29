package appmodel

import (
	"github.com/smalake/money-record-api/pkg/mysql"
)

type AppModel struct {
	MysqlCli *mysql.Client
}

func New(mysqlCli *mysql.Client) *AppModel {
	return &AppModel{MysqlCli: mysqlCli}
}
