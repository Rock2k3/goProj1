package main

import (
	app2 "app2"
	"goProject1/internal/pkg/httpServer"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		httpServer.Run(":8081")
	}()

	app2.TestApp2()

	wg.Wait()
}
