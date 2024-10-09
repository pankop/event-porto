package entity

import (
	"github.com/google/uuid"
	"github.com/pankop/event-porto/constants"
)

type Payments struct {
	Payment_ID         uuid.UUID               `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"payment_id"`
	Amount             int64                   `json:"amount"`
	Payment_Proof      string                  `json:"payment_proof"`
	Bank_Transfer_From string                  `json:"bank_transfer_from"`
	Name_Transfer_From string                  `json:"name_transfer_from"`
	Status             constants.PaymentStatus `json:"status"`
	Payment_Method     constants.PaymentMethod `json:"method"`
	Bank_ID            string                  `json:"bank_id"`
	Registrant_ID      uuid.UUID               `gorm:"type:uuid" json:"registrant_id"` // Ubah tipe data ke uuid.UUID
}
