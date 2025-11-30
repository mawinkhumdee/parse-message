package main

import (
	"context"
	"parse-message/config"
	"parse-message/db"
	"parse-message/handler"
	"parse-message/service"
	grpcServer "parse-message/transport/grpc"
	kafkaTransport "parse-message/transport/kafka"
	"parse-message/utils"
)

func main() {
	ctx := context.Background()
	cfg := config.LoadConfig()
	database := db.New(cfg.DB)
	util := utils.New(cfg)

	kafkaClients := kafkaTransport.NewClients(cfg)
	defer kafkaClients.Close()

	service := service.New(cfg, database, util, kafkaClients.ParseProducer, kafkaClients.UpdateProducer)

	go func() {
		if err := grpcServer.Start(50051, service); err != nil {
			panic(err)
		}
	}()

	handlers := handler.New(service)

	go func() {
		if err := kafkaClients.UpdateConsumer.Start(ctx, handlers.Update); err != nil {
			panic(err)
		}
	}()

	go func() {
		if err := kafkaClients.ParseConsumer.Start(ctx, handlers.Parse); err != nil {
			panic(err)
		}
	}()

	select {}
}
