package main

import (
	"root/src/core"
	"root/src/util"
)

func main() {
	util.LoadEnv()
	//config.InitRedis()
	//go consumer.StartConsumer()
	//go monitor.Init()
	core.StartServer()
}
