package interfaces

import "github.com/ashiqsabith123/notification-svc/pkg/domain"

type UserRepo interface {
	CreateNotification(data domain.Notifications) error
	GetAllNotification(query string) (notifications []domain.Notifications, err error)
	SaveFCMTokens(token domain.FcmTokens) error
	GetFCMToken(userID int) (token string, err error)
	UpdateNotificationStatus(commonID int, status string) error
}
