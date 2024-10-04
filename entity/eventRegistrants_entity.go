package entity

import (
	"time"

	"github.com/pankop/event-porto/constants"
)

type EventRegistrants struct {
	Regristant_ID     int64                        `gorm:"primary_key;column:regristant_id"`
	Status            constants.RegristationStatus `gorm:"default:pending" json:"status"`
	Reason            string                       `json:"reason"`
	Team_Name         string                       `json:"team_name"`
	School            string                       `json:"school"`
	Comp_Status       constants.CompStatus         `gorm:"default:penyisihan" json:"comp_status"`
	Regristation_Date time.Time                    `json:"regristation_date"`
	Account_ID        Account                      `gorm:"foreignKey:id;references:id"`
	Event_ID          []Events                     `gorm:"foreignKey:event_id;references:event_id"`
	Payment           []Payments                   `gorm:"foreignKey:payment_id;references:payment_id"`
	IoIDetailsID      *int64                       `gorm:"column:ioi_details_id"`
	IoIDetails        IoIDetails                   `gorm:"foreignKey:ioi_details_id;references:ioi_details_id;"`
	IMODetailsID      *int64                       `gorm:"column:imo_details_id"`
	IMODetails        IMODetails                   `gorm:"foreignKey:imo_details_id;references:imo_details_id"`
}
