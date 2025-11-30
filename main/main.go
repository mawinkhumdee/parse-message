package main

import (
	"context"
	"fmt"
	"parse-message/config"
	"parse-message/consumer"
	"parse-message/db"
	"parse-message/service"
	"parse-message/utils"
)

func main() {
	config := config.LoadConfig()
	db := db.New(config.DB)
	utils := utils.New(config)
	service := service.New(config, db, utils)

	ctx := context.Background()
	consumer := consumer.New(config)
	defer consumer.Close()

	handler := func(ctx context.Context, message string) error {
		fmt.Printf("Processing message: %s\n", message)
		result, err := service.ParseMessage(ctx, message)
		if err != nil {
			return fmt.Errorf("failed to parse message: %w", err)
		}
		fmt.Println("Parse Result:", result)

		insertedID, err := db.ParseResult.Insert(ctx, result)
		if err != nil {
			return fmt.Errorf("failed to insert result: %w", err)
		}
		fmt.Println("Inserted ID:", insertedID)
		return nil
	}

	if err := consumer.Start(ctx, handler); err != nil {
		panic(err)
	}
}
