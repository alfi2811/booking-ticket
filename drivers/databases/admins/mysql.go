package users

import (
	"booking-ticket/business/admins"
	"context"

	"gorm.io/gorm"
)

type MysqlAdminRepository struct {
	Conn *gorm.DB
}

func NewMysqlAdminRepository(conn *gorm.DB) admins.Repository {
	return &MysqlAdminRepository{
		Conn: conn,
	}
}

func (rep *MysqlAdminRepository) Login(ctx context.Context, email string, password string) (admins.Domain, error) {
	var admin Admins
	result := rep.Conn.First(&admin, "email = ? AND password = ?",
		email, password)

	if result.Error != nil {
		return admins.Domain{}, result.Error
	}

	return admin.ToDomain(), nil
}

func (rep *MysqlAdminRepository) Register(ctx context.Context, domain admins.Domain) (admins.Domain, error) {
	adminDB := FromDomain(domain)

	result := rep.Conn.Create(&adminDB)

	if result.Error != nil {
		return admins.Domain{}, result.Error
	}

	return adminDB.ToDomain(), nil
}

func (rep *MysqlAdminRepository) GetAdmin(ctx context.Context) ([]admins.Domain, error) {
	var user []Admins
	// var domainUser []users.Domain
	result := rep.Conn.Find(&user)

	if result.Error != nil {
		return []admins.Domain{}, result.Error
	}

	return ToListDomain(user), nil
}
