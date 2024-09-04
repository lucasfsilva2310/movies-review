package users

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}
