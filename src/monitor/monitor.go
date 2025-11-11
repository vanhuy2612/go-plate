package monitor

import (
	"fmt"
	"runtime"
	"time"
)

func printStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Alloc = %v MB \tTotalAlloc = %v MB \tSys = %v MB", m.Alloc/1024/1024, m.TotalAlloc/1024/1024, m.Sys/1024/1024)
	fmt.Printf("\tNumGC = %v \tNumCPU = %v \tNumGoroutine = %v", m.NumGC, runtime.NumCPU(), runtime.NumGoroutine())
	fmt.Printf("\n------------------------------------\n")
}

func Init()  {
	for {
		printStats()
		time.Sleep(10 * time.Second)
	}
}