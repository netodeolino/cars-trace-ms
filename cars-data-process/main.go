package main

import (
	"fmt"
	"time"

	"github.com/netodeolino/cars-trace-ms/cars-data-process/services/process"
)

var startTime time.Time

func uptime() time.Duration {
	return time.Since(startTime)
}

func init() {
	startTime = time.Now()
}

func main() {
	fmt.Println("Processing data...")

	process.Start()

	fmt.Printf("Process finished! Uptime %s\n", uptime())
}
