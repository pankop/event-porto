package entity

import (
	"github.com/pankop/event-porto/constants"
)

type AccountDetails struct {
	Account_ID   string            `gorm:"primary_key;column:account_id"`
	Name         string            `json:"name"`
	Phone_Number string            `json:"phone_number"`
	Jenjang      constants.Jenjang `json:"jenjang"`

	Timestamp
}
