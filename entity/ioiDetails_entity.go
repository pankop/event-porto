package entity

import (
	"github.com/google/uuid"
	"github.com/pankop/event-porto/constants"
)

type IoIDetails struct {
	IoIDetailsID       uuid.UUID          `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name               string             `json:"name"`
	Email              string             `json:"email"`
	PhoneNumber        string             `json:"phone_number"`
	Role               constants.Role_IOI `json:"role"`
	ImgIdentity        string             `json:"img_identity"`
	ImgFollowInstagram string             `json:"img_follow_instagram"`
	Link_Twibbon       string             `json:"link_twibbon"`
	Regristant_ID      []EventRegistrants `gorm:"foreignKey:regristant_id;references:regristant_id"`
}
