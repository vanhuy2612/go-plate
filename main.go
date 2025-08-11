package main

import (
	config "root/src/config"
	consumer "root/src/consumer"
	lib "root/src/lib"
	monitor "root/src/monitor"
)

func main() {
	config.Init()
	go consumer.StartConsumer()
	go monitor.Init()
	lib.StartServer()
}
