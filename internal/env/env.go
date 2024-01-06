package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type MysqlConfig struct {
	DBname   string
	Username string
	Password string
	Host     string
	Port     int
}

type MailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

var Mc MysqlConfig

func SetMysqlConfig() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	Mc.DBname = os.Getenv("MYSQL_NAME")
	Mc.Username = os.Getenv("MYSQL_USER")
	Mc.Password = os.Getenv("MYSQL_PASSWORD")
	Mc.Host = os.Getenv("MYSQL_HOST")
	Mc.Port, _ = strconv.Atoi(os.Getenv("MYSQL_PORT"))
	return nil
}

func GetMailConfig() (*MailConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return &MailConfig{}, err
	}
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	return &MailConfig{
		Host:     os.Getenv("MAIL_HOST"),
		Port:     port,
		Username: os.Getenv("MAIL_USERNAME"),
		Password: os.Getenv("MAIL_PASSWORD"),
		From:     os.Getenv("MAIL_FROM"),
	}, nil
}
