package auth

type AuthLoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthLoginAdmin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
