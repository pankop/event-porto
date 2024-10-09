package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/pankop/event-porto/constants"
)

type EventRegistrants struct {
	Registrant_ID     uuid.UUID                    `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Status            constants.RegristationStatus `gorm:"default:pending" json:"status"`
	Team_Name         string                       `json:"team_name"`
	School            string                       `json:"school"`
	Comp_Status       constants.CompStatus         `gorm:"default:penyisihan" json:"comp_status"`
	Registration_Date time.Time                    `json:"registration_date"`
	Account_ID        uuid.UUID                    `gorm:"type:uuid" json:"account_id"` // Sesuaikan tipe data
	Payment           []Payments                   `gorm:"foreignKey:Registrant_ID;references:Registrant_ID"`
	RegistrantDetails []RegistrationDetails        `gorm:"foreignKey:Registrant_ID;references:Registrant_ID"`
	Events            []Events                     `gorm:"many2many:event_registrants_events;foreignKey:Registrant_ID;joinForeignKey:Registrant_ID;References:Event_ID;joinReferences:Event_ID"`
}
