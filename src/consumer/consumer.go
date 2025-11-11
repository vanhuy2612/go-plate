package consumer

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"root/src/storage"
)

func StartConsumer() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"10.100.30.105:39091", "10.100.30.105:39092", "10.100.30.105:39093"},
		Topic:    "golang",
		GroupID:  "golang-group",
		MinBytes: 1, // 10KB
		MaxBytes: 10e6, // 10MB
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
		val, err := storage.Rdb.Get(storage.Ctx, "golang").Result()
		if err != nil {
			panic(err)
		}
		log.Printf("Redis value = %s", val)
	}
}
