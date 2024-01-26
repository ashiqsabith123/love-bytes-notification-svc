package service

import (
	"context"

	"github.com/ashiqsabith123/love-bytes-proto/notifications/pb"
	interfaces "github.com/ashiqsabith123/notification-svc/pkg/usecase/interface"
)

type UserService struct {
	UserUsecase interfaces.UserUsecase
	pb.UnimplementedNotificationServiceServer
}

func NewUserService(usecase interfaces.UserUsecase) UserService {
	return UserService{UserUsecase: usecase}
}

func (U *UserService) CreateNotification(ctx context.Context, req *pb.NotificationRequest) (*pb.NormalResponce, error) {
	err := U.UserUsecase.CreateNotification(req)

	if err != nil {
		return &pb.NormalResponce{}, nil
	}

	return &pb.NormalResponce{
		Message: "Notification created succesfully",
	}, nil
}
