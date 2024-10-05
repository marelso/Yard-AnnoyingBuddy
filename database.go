package main

import (
	"context"
	"firebase.google.com/go/db"
	"fmt"
)

func fetchDevices(ctx context.Context, client *db.Client) (DevicesResponse, error) {
	var devicesResponse DevicesResponse
	err := client.NewRef("/ESP32_Devices").Get(ctx, &devicesResponse.List)
	if err != nil {
		return DevicesResponse{}, fmt.Errorf("error reading from database: %w", err)
	}
	return devicesResponse, nil
}
