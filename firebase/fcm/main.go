package main

import (
	"io/ioutil"

	"google.golang.org/api/option"
	"github.com/acoshift/go-firebase-admin"
)

func main() {
	// Init App with service_account
	firApp, err := firebase.InitializeApp(context.Background(), firebase.AppOptions{
		ProjectID:      "YOUR_PROJECT_ID",
		DatabaseURL:    "YOUR_DATABASE_URL",
		APIKey:         "YOUR_API_KEY",
	}, option.WithCredentialsFile("service_account.json"))

	if err != nil {
		panic(err)
	}

	// FCM
	firFCM := firApp.FCM()

	// SendToDevice
	resp, err := firFCM.SendToDevice(context.Background(), "mydevicetoken",
		firebase.Message{Notification: firebase.Notification{
			Title: "Hello go firebase admin",
			Body:  "My little Big Notification",
			Color: "#ffcc33"},
		})

	if err != nil {
		panic(err)
	}

	// SendToDevices
	resp, err := firFCM.SendToDevices(context.Background(), []string{"mydevicetoken"},
		firebase.Message{Notification: firebase.Notification{
			Title: "Hello go firebase admin",
			Body:  "My little Big Notification",
			Color: "#ffcc33"},
		})

	if err != nil {
		panic(err)
	}

	// SubscribeDeviceToTopic
	resp, err := firFCM.SubscribeDeviceToTopic(context.Background(), "mydevicetoken", "/topics/gofirebaseadmin")
	// it's possible to ommit the "/topics/" prefix
	resp, err := firFCM.SubscribeDeviceToTopic(context.Background(), "mydevicetoken", "gofirebaseadmin")

	if err != nil {
		panic(err)
	}

	// UnSubscribeDeviceFromTopic
	resp, err := firFCM.UnSubscribeDeviceFromTopic(context.Background(), "mydevicetoken", "/topics/gofirebaseadmin")
	// it's possible to ommit the "/topics/" prefix
	resp, err := firFCM.UnSubscribeDeviceFromTopic(context.Background(), "mydevicetoken", "gofirebaseadmin")

	if err2 != nil {
		panic(err)
	}

}