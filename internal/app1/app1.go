package app1

import (
	"fmt"
	"os"
)

type AppSettings struct {
	Addr string
}

func Init() (AppSettings, error) {
	fmt.Println("Init App1")

	port := os.Getenv("HTTP_PORT_APP1")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)

	appSettings := AppSettings{Addr: addr}

	return appSettings, nil

}
