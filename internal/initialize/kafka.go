package initialize

import (
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/shinkaym/go-ecommerce-backend-api/global"
)

var KafkaProducer *kafka.Writer

func InitKafka() {
	global.KafkaProducer = &kafka.Writer{
		Addr:     kafka.TCP("localhost: 19092"),
		Topic:    "otp-auth-topic", // top name
		Balancer: &kafka.LeastBytes{},
	}
}
func CloseKafka() {
	if err := global.KafkaProducer.Close(); err != nil {
		log.Fatal("Failed to close kafka producer: %v", err)
	}
}