package main

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"fmt"
	"google.golang.org/api/option"
)

type Firebase struct {
	*db.Client
}

var database Firebase

func (db *Firebase) Connect() error {
	ctx := context.Background()
	opt := option.WithCredentialsFile("healthier-19122-firebase-adminsdk.json")
	config := &firebase.Config{DatabaseURL: "https://healthier-19122-default-rtdb.firebaseio.com/"}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		return fmt.Errorf("error initializing app: %v", err)
	}
	client, err := app.Database(ctx)
	if err != nil {
		return fmt.Errorf("error initializing database: %v", err)
	}
	db.Client = client
	return nil
}

func FirebaseDB() *Firebase {
	return &database
}
