package main

import (
	"root/src/consumer"
	"root/src/core"
	"root/src/monitor"
	"root/src/storage"
	"root/src/util"
)

func main() {
	util.LoadEnv()
	go storage.InitRedis()
	go consumer.StartConsumer()
	go monitor.Init()
	core.StartServer()
}
