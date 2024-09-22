//My Server

package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {

	go func() {
		fmt.Println("Started handling request")
		time.Sleep(5 * time.Second)
		fmt.Println("Finished handling request")
	}()

	fmt.Fprintf(w, "Request is being processed in a goroutine!\n")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
