package main

import (
	"context"
	"fmt"
	"parse-message/config"
	"parse-message/db"
	"parse-message/service"
	"parse-message/utils"
)

func main() {
	config := config.LoadConfig()
	db := db.New(config.DB)
	utils := utils.New()
	service := service.New(config, db, utils)

	ctx := context.Background()
	message := "I want to buy a new phone"
	result, err := service.ParseMessage(ctx, message)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
