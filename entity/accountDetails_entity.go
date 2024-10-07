package entity

import (
	"github.com/google/uuid"
	"github.com/pankop/event-porto/constants"
)

type AccountDetails struct {
	ID           uuid.UUID         `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name         string            `json:"name"`
	Phone_Number string            `json:"phone_number"`
	Jenjang      constants.Jenjang `json:"jenjang"`
	AccountID    string            `json:"account_id"`

	Timestamp
}
