package constants

type PaymentStatus string

const (
	Pending  PaymentStatus = "pending"
	Revision PaymentStatus = "revision"
	Success  PaymentStatus = "success"
	Rejected PaymentStatus = "rejected"
)
