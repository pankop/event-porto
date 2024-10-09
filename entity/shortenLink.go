package entity

type ShortenLink struct {
	Link_ID      int64  `gorm:"primary_key;auto_increment" json:"link_id"`
	ShortenLink  string `json:"shorten_link"`
	OriginalLink string `json:"original_link"`
	Account_ID   string `json:"account_id"`
}
