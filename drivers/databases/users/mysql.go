package users

import (
	"booking-ticket/business/users"
	"booking-ticket/controllers/users/requests"
	"context"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) users.Repository {
	return &MysqlUserRepository{
		Conn: conn,
	}
}

func (rep *MysqlUserRepository) Login(ctx context.Context, email string, password string) (users.Domain, error) {
	var user Users
	result := rep.Conn.First(&user, "email = ? AND password = ?",
		email, password)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil
}

func (rep *MysqlUserRepository) Register(ctx context.Context, userRegister requests.UserRegister) (users.Domain, error) {
	var user Users

	user.Email = userRegister.Email
	user.Password = userRegister.Password
	user.Fullname = userRegister.Fullname
	user.Gender = userRegister.Gender
	user.Phone = userRegister.Phone
	user.Status = 1

	result := rep.Conn.Create(&user)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil
}

func (rep *MysqlUserRepository) GetUser(ctx context.Context) ([]users.Domain, error) {
	var user []Users
	// var domainUser []users.Domain
	result := rep.Conn.Find(&user)

	if result.Error != nil {
		return []users.Domain{}, result.Error
	}

	return ToListDomain(user), nil
}
