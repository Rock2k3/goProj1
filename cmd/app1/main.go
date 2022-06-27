package main

import (
	"github.com/joho/godotenv"
	app1 "goProject1/internal/app1"
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

	appSettings, err := app1.Init()

	if err != nil {
		log.Fatal("failed to initialise")
	}

	go func() {
		httpServer.Run(appSettings.Addr)
	}()

	wg.Wait()
}
