package domain

// User holds user data
type User struct {
	ID       uint64 `json:"id"`
	FistName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
}
