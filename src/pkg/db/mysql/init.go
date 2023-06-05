package mysql

import (
	"github.com/Briofy/fs-go/src/pkg/db/database"
	"gorm.io/gorm"
)

func InitDB(cfg database.Config) (*gorm.DB, error) {
	//todo it should move to config
	var db *gorm.DB

	return db, nil
}
