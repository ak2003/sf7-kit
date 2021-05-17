package connections

import (
	"fmt"

	"gitlab.dataon.com/gophers/sf7-kit/shared/utils/config"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/jmoiron/sqlx"
)

func ConnSlave(logger log.Logger) (*sqlx.DB, error) {
	var err error
	var db *sqlx.DB
	{
		var (
			dbDriver = "sqlserverslave"
			dbUser   = config.GetDBUser(dbDriver)
			dbPass   = config.GetDBPass(dbDriver)
			dbHost   = config.GetDBHost(dbDriver)
			dbPort   = config.GetDBPort(dbDriver)
			dbName   = config.GetDBName(dbDriver)
		)
		var dbSource = fmt.Sprintf("server=%v;user id=%v;password=%v;port=%v;database=%v;encrypt=disable;", dbHost, dbUser, dbPass, dbPort, dbName)
		level.Info(logger).Log("dbInfo", dbSource)
		db, err = sqlx.Open("sqlserver", dbSource)
		return db, err
	}
}

func ConnMaster(logger log.Logger) (*sqlx.DB, error) {
	var err error
	var db *sqlx.DB
	{
		var (
			dbDriver = "sqlservermaster"
			dbUser   = config.GetDBUser(dbDriver)
			dbPass   = config.GetDBPass(dbDriver)
			dbHost   = config.GetDBHost(dbDriver)
			dbPort   = config.GetDBPort(dbDriver)
			dbName   = config.GetDBName(dbDriver)
		)
		var dbSource = fmt.Sprintf("server=%v;user id=%v;password=%v;port=%v;database=%v;encrypt=disable;", dbHost, dbUser, dbPass, dbPort, dbName)
		level.Info(logger).Log("dbInfo", dbSource)
		db, err = sqlx.Open("sqlserver", dbSource)
		return db, err
	}
}
