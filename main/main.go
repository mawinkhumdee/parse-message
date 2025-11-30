package main

import (
	"context"
	"fmt"
	"parse-message/config"
	"parse-message/consumer"
	"parse-message/db"
	"parse-message/model"
	"parse-message/producer"
	"parse-message/service"
	grpcServer "parse-message/transport/grpc"
	"parse-message/utils"
)

func main() {
	config := config.LoadConfig()
	db := db.New(config.DB)
	utils := utils.New(config)
	// Initialize Producer for parse-message topic
	parseMessageConfig := config.Kafka.Topics["parse-message"]
	parseProducer := producer.New(config.Kafka.Brokers, parseMessageConfig.Topic)
	defer parseProducer.Close()

	// Initialize Producer for update-message topic
	updateMessageConfig := config.Kafka.Topics["update-message"]
	updateProducer := producer.New(config.Kafka.Brokers, updateMessageConfig.Topic)
	defer updateProducer.Close()

	service := service.New(config, db, utils, parseProducer)

	ctx := context.Background()

	// Start gRPC Server in a goroutine
	go func() {
		if err := grpcServer.Start(50051, service); err != nil {
			panic(err)
		}
	}()

	// Initialize Consumer for parse-message topic
	parseConsumer := consumer.New(config.Kafka.Brokers, parseMessageConfig.Topic, parseMessageConfig.GroupID)
	defer parseConsumer.Close()

	handler := func(ctx context.Context, message string) error {
		fmt.Printf("Processing message: %s\n", message)
		// ... (comments omitted for brevity) ...

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

		// Produce to update-message topic
		updateMsg := fmt.Sprintf(`{"id": "TODO_ID", "status": "success"}`)
		if err := updateProducer.Produce(ctx, updateMsg); err != nil {
			fmt.Printf("Failed to produce update message: %v\n", err)
		}

		return nil
	}

	// Initialize Consumer for update-message topic
	updateConsumer := consumer.New(config.Kafka.Brokers, updateMessageConfig.Topic, updateMessageConfig.GroupID)
	defer updateConsumer.Close()

	updateHandler := func(ctx context.Context, message string) error {
		fmt.Printf("Processing update message: %s\n", message)
		// Parse message to get ID and Status
		// For now, just call service.UpdateMessage with dummy data to show wiring
		// In real impl, unmarshal JSON
		err := service.UpdateMessage(ctx, model.Message{ID: "TODO_ID", Status: "success"})
		if err != nil {
			fmt.Printf("Failed to update message: %v\n", err)
			return err
		}
		return nil
	}

	// Start Consumers
	go func() {
		if err := updateConsumer.Start(ctx, updateHandler); err != nil {
			panic(err)
		}
	}()

	// Start Consumer (blocking)
	if err := parseConsumer.Start(ctx, handler); err != nil {
		panic(err)
	}
}
