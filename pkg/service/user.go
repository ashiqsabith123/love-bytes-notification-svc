package service

import (
	"context"
	"net/http"

	"github.com/ashiqsabith123/love-bytes-proto/notifications/pb"
	interfaces "github.com/ashiqsabith123/notification-svc/pkg/usecase/interface"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type UserService struct {
	UserUsecase interfaces.UserUsecase
	pb.UnimplementedNotificationServiceServer
}

func NewUserService(usecase interfaces.UserUsecase) UserService {
	return UserService{UserUsecase: usecase}
}

func (U *UserService) CreateNotification(ctx context.Context, req *pb.Notification) (*pb.NormalResponce, error) {
	err := U.UserUsecase.CreateNotification(req)

	if err != nil {
		return &pb.NormalResponce{}, nil
	}

	return &pb.NormalResponce{
		Message: "Notification created succesfully",
	}, nil
}

func (U *UserService) GetAllNotifiacation(ctx context.Context, req *pb.GetNotificationRequest) (*pb.NotificationResponce, error) {

	notifications, err := U.UserUsecase.GetAllNotifications(req)

	if err != nil {
		return &pb.NotificationResponce{
			Code:    500,
			Message: "internal server error",
			Error: &anypb.Any{
				Value: []byte(err.Error()),
			},
		}, nil
	}

	data := make([]*pb.Notification, len(notifications))

	for i, v := range notifications {
		notificaion := &pb.Notification{
			CommonID: uint32(v.CommonID),
			SenderID: uint32(v.SenderID),
			Name:     v.Name,
			Image:    v.Image,
			Type:     v.Type,
			Status:   v.Status,
			Time:     v.CreatedAt.Format("03:04 PM, Mon 02 Jan 2006"),
		}

		data[i] = notificaion
	}

	allNotifications := pb.AllNotifications{
		Allnotification: data,
	}

	dataInBytes, err := proto.Marshal(&allNotifications)
	if err != nil {
		return &pb.NotificationResponce{
			Code:    http.StatusInternalServerError,
			Message: "Failed while marshaling",
			Error: &anypb.Any{
				Value: []byte(err.Error()),
			},
		}, nil
	}

	return &pb.NotificationResponce{
		Code:    http.StatusOK,
		Message: "Data fetched succesfully",
		Data: &anypb.Any{
			Value: dataInBytes,
		},
	}, nil

}
