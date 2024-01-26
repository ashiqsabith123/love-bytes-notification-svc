package repository

import (
	"github.com/ashiqsabith123/notification-svc/pkg/domain"
	interfaces "github.com/ashiqsabith123/notification-svc/pkg/repository/interface"
	"gorm.io/gorm"
)

type UserRepo struct {
	Postgres *gorm.DB
}

func NewUserRepo(db *gorm.DB) interfaces.UserRepo {
	return &UserRepo{Postgres: db}
}

func (U *UserRepo) CreateNotification(data domain.Notifications) error {

	err := U.Postgres.Create(&data).Error

	if err != nil {
		return err
	}

	return nil
}
