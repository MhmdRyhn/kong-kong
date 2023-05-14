package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got request in `/hello` endpoint.")
	io.WriteString(w, fmt.Sprintf("Hello, HTTP! @ %s", time.Now().Format("2006-01-02 03:04:05 PM")))
}

func main() {
	http.HandleFunc("/hello", getHello)
	port := 1100
	fmt.Println(fmt.Sprintf("Starting the server in port: %d", port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error starting the server: %s", err.Error()))
	}
}
