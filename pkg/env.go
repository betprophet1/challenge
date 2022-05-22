package pkg

import (
	"os"
	"reflect"
)

const (
	MYSQL_USER     = "MYSQL_USER"
	MYSQL_PASSWORD = "MYSQL_PASSWORD"
	MYSQL_DATABASE = "MYSQL_DATABASE"
	MYSQL_HOST     = "MYSQL_HOST"
	MYSQL_PORT     = "MYSQL_PORT"
)
type IEnv interface {
	Get() *Env
}

type Env struct {
	MysqlUser     string
	MysqlPassword string
	MysqlHost     string
	MysqlDatabase string
	MysqlPort     string
}

func Get() *Env {
	return &Env{
		MysqlUser:     os.Getenv(MYSQL_USER),
		MysqlPassword: os.Getenv(MYSQL_PASSWORD),
		MysqlHost:     os.Getenv(MYSQL_HOST),
		MysqlDatabase: os.Getenv(MYSQL_DATABASE),
		MysqlPort:     os.Getenv(MYSQL_PORT),
	}
}

func Default() *Env {
	return &Env{
		MysqlUser:     "wager",
		MysqlPassword: "123456",
		MysqlHost:     "localhost",
		MysqlDatabase: "wager",
		MysqlPort:     "3306",
	}
}

func (e *Env) IsEmpty() bool {
	return reflect.DeepEqual(e, &Env{})
}