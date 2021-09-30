package users

type UserUpdate struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Fullname string `json:"fullname" form:"fullname"`
	Gender   string `json:"gender" form:"gender"`
	Phone    string `json:"phone" form:"phone"`
}
