package admins

type AdminResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Token    string `json:"token"`
}
