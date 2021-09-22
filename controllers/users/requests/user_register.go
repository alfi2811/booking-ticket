package requests

type UserRegister struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
}
