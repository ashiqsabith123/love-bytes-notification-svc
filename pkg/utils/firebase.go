package pushnotification

import (
	"context"
	"encoding/json"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	logs "github.com/ashiqsabith123/love-bytes-proto/log"
	"github.com/ashiqsabith123/notification-svc/pkg/config"
	interfaces "github.com/ashiqsabith123/notification-svc/pkg/utils/interface"
	"google.golang.org/api/option"
)

type Utils struct {
	firebaseApp *firebase.App
}

func NewFirebaseApp(config config.Config) interfaces.Utils {

	firebaseCred, err := json.Marshal(config.FirebaseCred)

	if err != nil {
		fmt.Println(err)
	}

	opts := []option.ClientOption{option.WithCredentialsJSON(firebaseCred)}

	app, err := firebase.NewApp(context.Background(), nil, opts...)

	if err != nil {
		logs.ErrLog.Println("Error in initializing firebase app:", err)
		//return err
	}

	return &Utils{firebaseApp: app}
}

func (U *Utils) SendNotification(name, message, token, image string) {
	fcmClient, err := U.firebaseApp.Messaging(context.Background())

	if err != nil {
		logs.ErrLog.Println(err)
		return
	}

	_, err = fcmClient.Send(context.Background(), &messaging.Message{

		Android: &messaging.AndroidConfig{

			Notification: &messaging.AndroidNotification{
				Icon:  "https://love-bites-bucket.s3.us-east-2.amazonaws.com/" + image + ".jpeg",
				Color: "#FF5733",
				Title: name,
				Body:  message,
			},
		},

		Token: token,
	})

	if err != nil {
		logs.ErrLog.Println(err)
	}

}
