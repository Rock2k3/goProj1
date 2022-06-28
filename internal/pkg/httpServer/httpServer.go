package httpServer

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Run(addr string)  {
	http.HandleFunc("/health_check", healthCheck)

	srv := &http.Server{
		Addr: addr,
	}

	log.Printf("http server starting on port %s", strings.Split(addr, ":")[1])
	log.Fatal(srv.ListenAndServe())

}

func healthCheck (w http.ResponseWriter, r *http.Request)  {
	log.Println("healthCheck")
	_, err := fmt.Fprintf(w, "Ok")
	if err != nil {
		return
	}
}


