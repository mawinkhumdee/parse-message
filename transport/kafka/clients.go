package kafka

import (
	"log"
	"parse-message/config"
	"parse-message/consumer"
	"parse-message/producer"
)

type Clients struct {
	ParseProducer  producer.Producer
	UpdateProducer producer.Producer
	ParseConsumer  consumer.Consumer
	UpdateConsumer consumer.Consumer
}

func NewClients(cfg config.Config) Clients {
	parseTopic := cfg.Kafka.Topics["parse-message"]
	updateTopic := cfg.Kafka.Topics["update-message"]

	return Clients{
		ParseProducer:  producer.New(cfg.Kafka.Brokers, parseTopic.Topic),
		UpdateProducer: producer.New(cfg.Kafka.Brokers, updateTopic.Topic),
		ParseConsumer:  consumer.New(cfg.Kafka.Brokers, parseTopic.Topic, parseTopic.GroupID),
		UpdateConsumer: consumer.New(cfg.Kafka.Brokers, updateTopic.Topic, updateTopic.GroupID),
	}
}

func (c Clients) Close() {
	if err := c.ParseProducer.Close(); err != nil {
		log.Printf("failed to close parse producer: %v", err)
	}
	if err := c.UpdateProducer.Close(); err != nil {
		log.Printf("failed to close update producer: %v", err)
	}
	if err := c.ParseConsumer.Close(); err != nil {
		log.Printf("failed to close parse consumer: %v", err)
	}
	if err := c.UpdateConsumer.Close(); err != nil {
		log.Printf("failed to close update consumer: %v", err)
	}
}
