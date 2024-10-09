package entity

import (
	"time"

	"github.com/google/uuid"
)

type Events struct {
	Event_ID    uuid.UUID          `gorm:"type:uuid;primary_key;default:uuid_generate_v4();column:event_id" json:"event_id"`
	Name        string             `json:"name"`
	Quota       int                `json:"quota"`
	Price       int                `json:"price"`
	Start_Date  time.Time          `json:"start_date"`
	End_Date    time.Time          `json:"end_date"`
	Registrants []EventRegistrants `gorm:"many2many:event_registrants_events;foreignKey:Event_ID;joinForeignKey:Event_ID;References:Registrant_ID;joinReferences:Registrant_ID"`
}
