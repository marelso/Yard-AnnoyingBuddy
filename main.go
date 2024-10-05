package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	err := connect()
	if err != nil {
		log.Println(err)
		return
	}

	for {
		var wg sync.WaitGroup
		data, err := fetchDevices(context.Background(), fb.database)
		if err != nil {
			log.Fatalf("Failed to fetch entries: %v\n", err)
		}

		for id, device := range data.List {
			wg.Add(1)
			go func(deviceID string, d Device) {
				defer wg.Done()
				processDevice(deviceID, d)
			}(id, device)
		}

		wg.Wait()
		fmt.Println("All devices processed.")

		time.Sleep(60 * time.Second)
	}
}

func processDevice(id string, device Device) {
	fmt.Printf("Now processing device : %s\n", id)
	SendPush("Title", "Message body", "cT9UlTL-TgaHjBH518zyVH:APA91bE85WIpjriKUSWolNwjAXySykoTXhlJ_uTsTxszssAvjqPWmRB9seB9ijGbXl7K0F5FPHScE7PynAqG6FVmzobbtZDAMnQl_TLmrfmPsMcNZmlB-w5OBURs5ro_lSXnqMmMHDYO")
}
