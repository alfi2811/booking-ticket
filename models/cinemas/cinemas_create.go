package cinemas

type CinemaCreate struct {
	Name     string `json:"name" form:"name" gorm:"unique"`
	Location string `json:"location" form:"location"`
	Maps     string `json:"maps" form:"maps"`
	Phone    string `json:"phone" form:"phone"`
}
