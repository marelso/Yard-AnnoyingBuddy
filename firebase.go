package main

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"firebase.google.com/go/messaging"
	"fmt"
	"google.golang.org/api/option"
)

type Firebase struct {
	database  *db.Client
	messaging *messaging.Client
}

var fb Firebase

func connect() error {
	ctx := context.Background()
	opt := option.WithCredentialsFile("credentials.json")
	config := &firebase.Config{DatabaseURL: "https://healthier-19122-default-rtdb.firebaseio.com/"}

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		return fmt.Errorf("error initializing app: %v", err)
	}

	dbClient, err := app.Database(ctx)
	if err != nil {
		return fmt.Errorf("error initializing database: %v", err)
	}
	fb.database = dbClient

	messageClient, err := app.Messaging(context.Background())
	if err != nil {
		return err
	}
	fb.messaging = messageClient

	return nil
}
