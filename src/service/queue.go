package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/twmb/franz-go/pkg/kgo"
	"root/src/producer"
	"root/src/util"
	"sync"
)

type IQueueService interface {
	Publish2Queue()
}

type QueueService struct {

}

func (q QueueService) Publish2Queue(c *gin.Context)  {
	ctx := context.Background()
	var wg sync.WaitGroup
	wg.Add(1)
	topic := util.GetEnv("KAFKA_TOPIC")
	record := &kgo.Record{Topic: topic, Value: []byte("bar")}
	producer.Client.Produce(ctx, record, func(_ *kgo.Record, err error) {
		defer wg.Done()
		if err != nil {
			fmt.Printf("record had a produce error: %v\n", err)
		}

	})
	wg.Wait()
}

