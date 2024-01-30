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

	fmt.Println("https://love-bites-bucket.s3.us-east-2.amazonaws.com/" + image + ".jpeg")

	if err != nil {
		logs.ErrLog.Fatal(err)
	}

	response, err := fcmClient.Send(context.Background(), &messaging.Message{

		Android: &messaging.AndroidConfig{

			Notification: &messaging.AndroidNotification{

				Title:    name,
				Body:     message,
				ImageURL: "https://love-bites-bucket.s3.us-east-2.amazonaws.com/" + image + ".jpeg",
			},
		},

		// Notification: &messaging.Notification{

		// 	Title:    name,
		// 	Body:     message,
		// 	ImageURL: "https://love-bites-bucket.s3.us-east-2.amazonaws.com/" + image + ".jpeg",
		// },
		Token: "f-AahZaTTrmJpGfQa08HOZ:APA91bFZAOIwzZdBrdi5inx9wu5r6IZ-z7ahfygXqMA7Xy5glAU2f0lpGs3Uuigtq3powncf0KHlVPHXMvVAKbERKaip-gOOS72k4-hP1-jICSyRJVe-ozvla42aQJuP4W-SIyRydJlC", // it's a single device token
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response)
}
