package entity

import (
	"time"
)

type Events struct {
	Event_ID         string             `gorm:"primary_key;column:event_id"`
	Name             string             `json:"name"`
	Quota            int                `json:"quota"`
	Price            int                `json:"price"`
	Start_Date       time.Time          `json:"start_date"`
	End_Date         time.Time          `json:"end_date"`
	EventRegistrants []EventRegistrants `gorm:"foreignKey:regristant_id;references:regristant_idegristant_ID"`
}
