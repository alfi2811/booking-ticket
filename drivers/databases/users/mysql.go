package users

import (
	"booking-ticket/business/users"
	"context"
	"fmt"

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
	result := rep.Conn.First(&user, "email = ? ", email)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil
}

func (rep *MysqlUserRepository) Register(ctx context.Context, domain users.Domain) (users.Domain, error) {
	user := FromDomain(domain)

	result := rep.Conn.Create(&user)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil
}

func (rep *MysqlUserRepository) GetUser(ctx context.Context) ([]users.Domain, error) {
	var user []Users
	// var domainUser []users.Domain
	result := rep.Conn.Debug().Find(&user)

	if result.Error != nil {
		return []users.Domain{}, result.Error
	}

	return ToListDomain(user), nil
}

func (rep *MysqlUserRepository) ListUserBooking(ctx context.Context) ([]users.Domain, error) {
	var user []Users
	// var domainUser []users.Domain
	result := rep.Conn.Debug().Preload("Booking").Find(&user)
	fmt.Println(user)
	if result.Error != nil {
		return []users.Domain{}, result.Error
	}

	return ToListDomain(user), nil
}
