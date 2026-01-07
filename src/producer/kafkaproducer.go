package producer

import (
	"github.com/twmb/franz-go/pkg/kgo"
	"root/src/util"
)
var Client *kgo.Client

type IKafkaProducer interface {
	Init()
}

type KafkaProducer struct {

}

func Init() {
	seeds := util.GetEnvWithComma("KAFKA_BROKER", ",")
	cl, err := kgo.NewClient(
		kgo.SeedBrokers(seeds...),
		kgo.ConsumerGroup("my-group-identifier"),
		kgo.ConsumeTopics("foo"),
	)
	if err != nil {
		panic(err)
	}
	defer cl.Close()
	Client = cl
}



