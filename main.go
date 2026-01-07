package main

import (
	"root/src/consumer"
	"root/src/core"
	"root/src/producer"
	"root/src/util"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	util.LoadEnv()
	go consumer.StartConsumer()
	go producer.Init()
	//go monitor.Init()
	core.StartServer()
}
