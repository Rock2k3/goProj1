package main

import (
	app2 "app2"
	"app2/handlers"
	"github.com/joho/godotenv"
	"goProject1/internal/pkg/httpServer"
	"log"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	go func() {
		handlers.AddHandlers()
		httpServer.Run(":8081")
	}()

	app2.TestApp2()

	wg.Wait()
}
