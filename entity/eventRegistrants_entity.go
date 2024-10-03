package entity

import (
	"time"

	"github.com/pankop/event-porto/constants"
)

type EventRegistrants struct {
	Regristant_ID     int64                        `gorm:"primary_key;column:regristant_id"`
	Status            constants.RegristationStatus `json:"status"`
	Reason            string                       `json:"reason"`
	Team_Name         string                       `json:"team_name"`
	School            string                       `json:"school"`
	Comp_Status       constants.CompStatus         `json:"comp_status"`
	Regristation_Date time.Time                    `json:"regristation_date"`
	ID                Account                      `gorm:"foreignKey:regristant_id;references:regristant_ID"`
	Event_ID          []Events                     `gorm:"foreignKey:event_id;references:event_id"`
	Payment           []Payments                   `gorm:"foreignKey:payment_id;references:payment_id"`
}
