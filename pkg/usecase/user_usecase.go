package usecase

import (
	"fmt"
	"sort"

	"github.com/ashiqsabith123/love-bytes-proto/notifications/pb"
	"github.com/ashiqsabith123/notification-svc/pkg/domain"
	repo "github.com/ashiqsabith123/notification-svc/pkg/repository/interface"
	usecase "github.com/ashiqsabith123/notification-svc/pkg/usecase/interface"
	utils "github.com/ashiqsabith123/notification-svc/pkg/utils/interface"
	"github.com/jinzhu/copier"
)

type UserUsecase struct {
	UserRepo repo.UserRepo
	Utils    utils.Utils
}

func NewUserUsecase(repo repo.UserRepo, Utils utils.Utils) usecase.UserUsecase {
	return &UserUsecase{UserRepo: repo, Utils: Utils}
}

func (U *UserUsecase) CreateNotification(req *pb.Notification) error {

	var notification domain.Notifications

	copier.Copy(&notification, &req)
	U.Utils.SendNotification(notification.Name, "Send you an interest request", "", notification.Image)

	err := U.UserRepo.CreateNotification(notification)

	if err != nil {
		return err
	}

	return nil

}

func (U *UserUsecase) GetAllNotifications(req *pb.GetNotificationRequest) ([]domain.Notifications, error) {

	query := "SELECT * FROM notifications WHERE receiver_id=" + fmt.Sprint(req.UserID)

	if req.Type != "" {
		query += " AND type='" + fmt.Sprint(req.Type) + "'"
	}

	if req.Day != "" {
		query += " AND created_at >= CURRENT_DATE - INTERVAL '" + req.Day + " days'"
	}

	notifications, err := U.UserRepo.GetAllNotification(query)

	if err != nil {
		return []domain.Notifications{}, err
	}

	sort.Slice(notifications, func(i, j int) bool {
		return notifications[i].CreatedAt.After(notifications[j].CreatedAt)
	})

	return notifications, nil

}

func (U *UserUsecase) SaveFCMTokens() error {

	var token domain.FcmTokens
	if err := U.UserRepo.SaveFCMTokens(token); err != nil {
		return err
	}

	return nil
}
