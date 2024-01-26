package usecase

import (
	"github.com/ashiqsabith123/love-bytes-proto/notifications/pb"
	"github.com/ashiqsabith123/notification-svc/pkg/domain"
	repo "github.com/ashiqsabith123/notification-svc/pkg/repository/interface"
	usecase "github.com/ashiqsabith123/notification-svc/pkg/usecase/interface"
	"github.com/jinzhu/copier"
)

type UserUsecase struct {
	UserRepo repo.UserRepo
}

func NewUserUsecase(repo repo.UserRepo) usecase.UserUsecase {
	return &UserUsecase{UserRepo: repo}
}

func (U *UserUsecase) CreateNotification(req *pb.NotificationRequest) error {

	var notification domain.Notifications

	copier.Copy(&notification, &req)

	err := U.UserRepo.CreateNotification(notification)

	if err != nil {
		return err
	}

	return nil

}
