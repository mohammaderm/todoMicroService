package app

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/mohammaderm/todoMicroService/todoService/config"
	"github.com/mohammaderm/todoMicroService/todoService/pkg/logger"
)

func DBconnection(logger logger.Logger, config config.Database) (*sqlx.DB, func(), error) {
	con, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", config.Mysql.Username, config.Mysql.Password, config.Mysql.Host, config.Mysql.Port, config.Mysql.Database))
	if err != nil {
		return nil, func() {}, err
	}
	return con, func() {
		if err := con.Close(); err != nil {
			logger.Warning("failed to close db connection", map[string]interface{}{
				"error": err.Error(),
			})
		}
	}, nil

}
