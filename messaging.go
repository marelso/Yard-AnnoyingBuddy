package main

import (
	"context"
	"firebase.google.com/go/messaging"
	"fmt"
	"log"
)

func SendPush(title string, body string, token string) {
	message := &messaging.Message{
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
		},
		Notification: &messaging.Notification{
			Title:    title,
			Body:     body,
			ImageURL: "https://image.lexica.art/md2_webp/3376abb8-a0e1-46a9-8f20-73d369c07730",
		},
		Token: token,
	}

	response, err := fb.messaging.Send(context.Background(), message)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Successfully sent message:", response)
}
