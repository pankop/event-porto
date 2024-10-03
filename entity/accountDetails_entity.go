package entity

import (
	"github.com/Caknoooo/go-gin-clean-starter/constants"
)

type AccountDetails struct {
	Account_ID   string  `gorm:"primary_key;column:account_id"`
	Name         string  `json:"name"`
	Phone_Number string  `json:"phone_number"`
	Province     string  `json:"province"`
	Jenjang      constants.Jenjang `json:"jenjang"`

	Timestamp
}
