package pkg

import (
	"github.com/joho/godotenv"
)
type IEnv interface {
	
}
type Env struct {
	MysqlUsername string `json:"mysql_username"`
	MysqlPassword string `json:"mysql_password"`
	MysqlHost     string `json:"mysql_host"`
	MysqlDatabase string `json:"mysql_database"`

}

func (e *Env) Load()  {
	err := godotenv.Load()
	if err != nil {
		return
	}
}