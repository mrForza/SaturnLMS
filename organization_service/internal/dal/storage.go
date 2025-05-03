package dal

import (
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mrForza/SaturnLMS/organization_service/internal/configs"
)

var Db *sqlx.DB

func InitDB() error {
	var cfg = *configs.DbConfig
	if dbConfig, ok := cfg.(configs.DatabaseConfig); ok {
		var err error
		Db, err = sqlx.Connect("postgres", dbConfig.ToDataSourceName())
		if err != nil {
			log.Printf("no connection to postgres db: %s", err.Error())
			log.Printf("no connection to postgres db: %s", dbConfig.ToDataSourceName())
			log.Fatalf("no connection to postgres db: %s", err.Error())
		}
		return Db.Ping()
	}
	return errors.New("incorrect DatabaseConfig")
}
