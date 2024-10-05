package main

import (
	"context"
	"firebase.google.com/go/db"
	"fmt"
	"yard-AnnoyingBuddy/domain"
)

func fetchDevices(ctx context.Context, client *db.Client) (map[string]domain.Device, error) {
	var devicesResponse map[string]domain.Device
	err := client.NewRef("/ESP32_Devices").Get(ctx, &devicesResponse)
	if err != nil {
		return map[string]domain.Device{}, fmt.Errorf("error reading from database: %w", err)
	}
	return devicesResponse, nil
}
