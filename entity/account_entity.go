package entity

import (
	"github.com/google/uuid"
	"github.com/pankop/event-porto/helpers"
	"gorm.io/gorm"
)

type Role string
type PaymentStatus string
type Role_IOI string

const (
	Leader Role_IOI = "leader"
	Member Role_IOI = "member"
)

type Account struct {
	ID              uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Email           string         `json:"email"`
	Password        string         `json:"password"`
	Role            string         `json:"role"`
	IsEmailVerified bool           `json:"is_verified"`
	AccountDetails  AccountDetails `gorm:"foreignKey:Account_ID"`

	Timestamp
}

func (u *Account) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var err error
	// u.ID = uuid.New()
	u.Password, err = helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}
