package db

import (
	"errors"
	"github.com/Briofy/fs-go/src/pkg/db/database"
	"github.com/Briofy/fs-go/src/pkg/db/mysql"
	"github.com/Briofy/fs-go/src/pkg/db/postgres"
	"gorm.io/gorm"
)

type InitDB func(cfg database.Config) (*gorm.DB, error)

func GetInitDb(databaseDriver database.Driver) (InitDB, error) {
	switch databaseDriver {
	case database.MysqlDatabaseDriver:
		return mysql.InitDB, nil
	case database.PostgresDatabaseDriver:
		return postgres.InitDB, nil
	default:
		return nil, errors.New("unknown driver")

	}
}
