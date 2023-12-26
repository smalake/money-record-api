package env

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Environment struct {
	Mysql MysqlConfig
}

type MysqlConfig struct {
	DBname   string
	Username string
	Password string
	Host     string
	Port     int
}

func Load() Environment {
	var e Environment
	e.Mysql = setMysqlConfig()
	return e
}

func setMysqlConfig() MysqlConfig {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}
	var mc MysqlConfig
	mc.DBname = os.Getenv("MYSQL_NAME")
	mc.Username = os.Getenv("MYSQL_USER")
	mc.Password = os.Getenv("MYSQL_PASSWORD")
	mc.Host = os.Getenv("MYSQL_HOST")
	mc.Port, _ = strconv.Atoi(os.Getenv("MYSQL_PORT"))
	return mc
}
