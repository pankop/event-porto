package entity

import (
	"github.com/pankop/event-porto/constants"
)

type RegistrationDetails struct {
	RegistrationDetailsID string             `gorm:"primary_key;default:uuid_generate_v4()" json:"id"`
	Name                  string             `json:"name"`
	Email                 string             `json:"email"`
	PhoneNumber           string             `json:"phone_number"`
	Role                  constants.TeamRole `json:"role"`
	ImgIdentity           string             `json:"img_identity"`
	ImgFollowInstagram    string             `json:"img_follow_instagram"`
	Link_Twibbon          string             `json:"link_twibbon"`
	Registrant_ID         []EventRegistrants `gorm:"foreignKey:registrant_id;references:registrant_id"`
}
