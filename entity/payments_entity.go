package entity

import (
	"github.com/Caknoooo/go-gin-clean-starter/constants"
)

type Payments struct {
	Payment_ID         int64                   `gorm:"primary_key;column:payment_id"`
	Regristant_ID      int64                   `json:"regristant_id"`
	Amount             int64                   `json:"amount"`
	Payment_Proof      string                  `json:"payment_proof"`
	Bank_Transfer_From string                  `json:"bank_transfer_from"`
	Name_Transfer_From string                  `json:"name_transfer_from"`
	Status             constants.PaymentStatus `json:"status"`
}
