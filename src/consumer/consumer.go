package consumer

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"root/src/util"
)

func StartConsumer() {
	var addrs []string
	addrs = util.GetEnvWithComma("KAFKA_BROKER", ",")
	topic := util.GetEnv("KAFKA_TOPIC")
	groupId := util.GetEnv("KAFKA_GROUPID")
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        addrs,
		Topic:          topic,
		GroupID:        groupId,
		MinBytes:       1,    // 10KB
		MaxBytes:       10e6, // 10MB
		StartOffset:    kafka.FirstOffset,
		CommitInterval: 0,
	})
	defer reader.Close()

	log.Printf("************** Kafka consumer is running ****************")
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("failed to read message:", err)
		}
		log.Printf("Received message at offset %d: key=%s value=%s\n",
			msg.Offset, string(msg.Key), string(msg.Value))
	}
}
