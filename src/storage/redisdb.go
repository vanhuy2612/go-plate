package storage

import (
	"context"
	"log"
	"root/src/util"

	"github.com/redis/go-redis/v9"
)

var (
	Rdb *redis.ClusterClient
	Ctx = context.Background()
)

// Init connects to the Redis cluster
func InitRedis() {
	var addrs []string
	var password string
	addrs = util.GetEnvWithComma("REDIS_HOST", ",")
	password = util.GetEnv("REDIS_PASS")
	Rdb = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    addrs,
		Password: password,
	})

	if err := Rdb.Ping(Ctx).Err(); err != nil {
		log.Fatalf("Could not connect to Redis Cluster: %v %v", addrs, err)
	}

	log.Println("Connected to Redis Cluster")
}
