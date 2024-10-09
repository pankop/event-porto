package migrations

import (
	"github.com/pankop/event-porto/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entity.Account{},
		&entity.AccountDetails{},
		&entity.BankList{},
		&entity.Payments{},
		&entity.EventRegistrants{},
		&entity.RegistrationDetails{},
		&entity.Events{},
		&entity.ShortenLink{},
	); err != nil {
		return err
	}

	return nil
}
