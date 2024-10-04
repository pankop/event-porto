package entity

import (
	"github.com/pankop/event-porto/constants"
)

type Payments struct {
	Payment_ID         int64                   `gorm:"primary_key;column:payment_id"`
	Amount             int64                   `json:"amount"`
	Payment_Proof      string                  `json:"payment_proof"`
	Bank_Transfer_From string                  `json:"bank_transfer_from"`
	Name_Transfer_From string                  `json:"name_transfer_from"`
	Status             constants.PaymentStatus `json:"status"`
	Method             constants.PaymentMethod `json:"method"`
	BankList           BankList                `gorm:"foreignKey:Bank_ID;references:Bank_ID"`
	registrant_ID      []EventRegistrants      `gorm:"foreignKey:registrant_id;references:registrant_id"`
}
