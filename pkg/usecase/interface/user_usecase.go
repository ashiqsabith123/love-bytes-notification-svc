package interfaces

import (
	"github.com/ashiqsabith123/love-bytes-proto/notifications/pb"
	"github.com/ashiqsabith123/notification-svc/pkg/domain"
)

type UserUsecase interface {
	CreateNotification(req *pb.Notification) error
	GetAllNotifications(req *pb.GetNotificationRequest) ([]domain.Notifications, error)
}
