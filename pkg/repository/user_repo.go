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

func (U *UserRepo) GetAllNotification(query string) (notifications []domain.Notifications, err error) {

	if err = U.Postgres.Raw(query).Scan(&notifications).Error; err != nil {
		return notifications, err
	}

	return notifications, nil
}

func (U *UserRepo) SaveFCMTokens(token domain.FcmTokens) error {
	
	if err := U.Postgres.Create(&token).Error; err != nil {
		return err
	}

	return nil
}
