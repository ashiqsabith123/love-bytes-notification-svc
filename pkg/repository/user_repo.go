package repository

import (
	"fmt"

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

func (U *UserRepo) GetFCMToken(userID int) (token string, err error) {

	query := "SELECT token FROM fcm_tokens WHERE user_id = $1"

	err = U.Postgres.Raw(query, userID).Scan(&token).Error

	if err != nil {
		return "", err
	}

	return token, nil

}

func (U *UserRepo) UpdateNotificationStatus(commonID int, status string) error {
	query := "UPDATE notifications SET status= $1 WHERE common_id = $2;"

	fmt.Println("hrewwwwwwwwwwwww")

	err := U.Postgres.Exec(query, status, commonID).Error

	if err != nil {
		return err
	}

	return nil
}
