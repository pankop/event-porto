package entity

type BankList struct {
	Bank_ID        string     `gorm:"primary_key;column:bank_id"`
	Account_Name  string     `json:"account_name"`
	Account_Number string     `json:"account_number"`
	Bank_Name      string     `json:"bank_name"`
	Payment        []Payments `gorm:"foreignKey:bank_id;references:bank_id"`
}
