package migrations

import (
	"github.com/pankop/event-porto/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entity.User{},
	); err != nil {
		return err
	}

	return nil
}
