package user

type Role string

const (
	Zeus    Role = "admin"
	Hunter  Role = "hunter"
	Client  Role = "client"
	Scooter Role = "scooter"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password" db:"passwd"`
	Role     Role   `json:"role"`
}
