package entity

type BankList struct {
	Bank_ID        string `gorm:"primary_key;column:bank_id"`
	Account_Naame  string `json:"account_name"`
	Account_Number string `json:"account_number"`
	Bank_Name      string `json:"bank_name"`
}
