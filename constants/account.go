package constants

type Role string
type Jenjang string

const (
	User  Role = "user"
	Admin Role = "admin"
)

const (
	Mahsiswa Jenjang = "mahasiswa"
	Pelajar  Jenjang = "pelajar"
)
