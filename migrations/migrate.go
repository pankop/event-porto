package migrations

import (
	"github.com/pankop/event-porto/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entity.Account{},
		&entity.AccountDetails{},
	); err != nil {
		return err
	}

	return nil
}
