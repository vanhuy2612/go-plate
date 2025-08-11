package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	Rdb *redis.ClusterClient
	Ctx = context.Background()
)

// Init connects to the Redis cluster
func Init() {
	Rdb = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"10.100.30.105:7001",
			"10.100.30.105:7002",
			"10.100.30.105:7003",
			"10.100.30.105:7004",
			"10.100.30.105:7005",
			"10.100.30.105:7006",
		},
		Password: "kbsv",
	})

	if err := Rdb.Ping(Ctx).Err(); err != nil {
		log.Fatalf("Could not connect to Redis Cluster: %v", err)
	}

	log.Println("Connected to Redis Cluster")
}