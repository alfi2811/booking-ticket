package auth

type AuthLoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
