package entity

import (
	"time"

	"github.com/pankop/event-porto/constants"
)

type EventRegistrants struct {
	Registrant_ID     int64                        `gorm:"primary_key;column:registrant_id"`
	Status            constants.RegristationStatus `gorm:"default:pending" json:"status"`
	Team_Name         string                       `json:"team_name"`
	School            string                       `json:"school"`
	Comp_Status       constants.CompStatus         `gorm:"default:penyisihan" json:"comp_status"`
	Registration_Date time.Time                    `json:"registration_date"`
	Account_ID        Account                      `gorm:"foreignKey:id;references:id"`
	Event_ID          []Events                     `gorm:"foreignKey:event_id;references:event_id"`
	Payment           []Payments                   `gorm:"foreignKey:payment_id;references:payment_id"`
	RegistrantDetails []RegistrationDetails        `gorm:"foreignKey:registrant_id;references:registrant_id"`
}
