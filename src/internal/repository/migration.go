package repository

import "gorm.io/gorm"

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&Attach{},
		&Attachable{},
	)
	if err != nil {
		return err
	}
	return nil
}
