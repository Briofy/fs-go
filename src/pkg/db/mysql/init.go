package mysql

import (
	"github.com/Briofy/fs-go/src/pkg/db/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg database.Config) (*gorm.DB, error) {
	//todo it should move to config
	dsn := cfg.GetDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}
