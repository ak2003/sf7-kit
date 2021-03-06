package database

import (
	"database/sql"
	"os"

	"gitlab.dataon.com/gophers/sf7-kit/shared/utils/config"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type PgSqlDB struct {
	DB *sql.DB
}

func NewDB(logger log.Logger) *sql.DB {
	var (
		dbDriver = "postgresql"
		dbUser   = config.GetDBUser(dbDriver)
		dbPass   = config.GetDBPass(dbDriver)
		dbHost   = config.GetDBHost(dbDriver)
		dbPort   = config.GetDBPort(dbDriver)
		dbName   = config.GetDBName(dbDriver)
	)
	var dbSource = "postgresql://" + dbUser + ":" + dbPass + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable"
	level.Info(logger).Log("dbInfo", dbSource)
	db, err := sql.Open("postgres", dbSource)
	if err != nil {
		level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}

	return db
}
