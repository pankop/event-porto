package constants

type RegristationStatus string
type CompStatus string

const (
	RegristationStatusPending RegristationStatus = "pending"
	RegristationStatusActive  RegristationStatus = "revision"
	RegristationStatusSuccess RegristationStatus = "success"
	RegristationStatusFailed  RegristationStatus = "failed"
)

const (
	CompStatusPenyisihan CompStatus = "penyisihan"
	CompStatusSemifinal  CompStatus = "semifinal"
	CompStatusFinal      CompStatus = "final"
)
