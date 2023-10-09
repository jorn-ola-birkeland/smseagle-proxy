package main

import (
	"errors"
	"fmt"
	"kartverket.no/smseagle-proxy/alerter"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/webhook", alerter.HandleWebhook)

	err := http.ListenAndServe(":8080", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
