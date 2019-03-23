package user

type Role string

const (
	Zeus   Role = "zeus"
	Hunter Role = "hunter"
	Client Role = "client"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Pass     string `json:"pass" db:"passwd"`
	Role     Role   `json:"role"`
}
