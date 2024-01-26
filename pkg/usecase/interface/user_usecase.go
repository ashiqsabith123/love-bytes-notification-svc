package interfaces

import "github.com/ashiqsabith123/love-bytes-proto/notifications/pb"

type UserUsecase interface {
	CreateNotification(req *pb.NotificationRequest) error
}
