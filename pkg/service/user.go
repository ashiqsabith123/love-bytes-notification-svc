package service

import (
	"github.com/ashiqsabith123/love-bytes-proto/notifications/pb"
	"github.com/ashiqsabith123/notification-svc/pkg/usecase"
)

type UserService struct {
	UserUsecase usecase.UserUsecase
	pb.UnimplementedNotificationServiceServer
}

func NewUserService(usecase usecase.UserUsecase) UserService {
	return UserService{UserUsecase: usecase}
}
