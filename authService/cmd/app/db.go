package app

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mohammaderm/authService/config"
	"github.com/mohammaderm/authService/pkg/logger"
)

func DBconnection(logger logger.Logger, config *config.Database) (*sqlx.DB, func(), error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Postgresql.Host,
		config.Postgresql.Port,
		config.Postgresql.Username,
		config.Postgresql.Password,
		config.Postgresql.Database,
	)
	con, err := sqlx.Connect("postgres", dsn)
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
