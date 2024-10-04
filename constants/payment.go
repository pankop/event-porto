package constants

type PaymentStatus string
type PaymentMethod string

const (
	Pending  PaymentStatus = "pending"
	Revision PaymentStatus = "revision"
	Success  PaymentStatus = "success"
	Rejected PaymentStatus = "rejected"
)

const (
	Transfer PaymentMethod = "transfer"
	Cash     PaymentMethod = "cash"
)
