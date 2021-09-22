package users

import "time"

type User struct {
	ID        int       `json:"id" form:"id" gorm:"primaryKey" `
	Email     string    `json:"email" form:"email" gorm:"unique"`
	Password  string    `json:"password" form:"password"`
	Fullname  string    `json:"fullname" form:"fullname"`
	Gender    string    `json:"gender" form:"gender"`
	Phone     string    `json:"phone" form:"phone"`
	Status    int       `json:"status" form:"status"`
	CreatedAt time.Time `json:"createdAt" form:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt"`
}
