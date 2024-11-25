package server

import (
	"os"
	"fmt"
	"net/http"
)

func Run() {
	addr := fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	router := buildRouter()
	fmt.Printf("Launching server on: %s\n", addr)
	err := http.ListenAndServe(addr, router)
	fmt.Printf("Server shutting down. Error: %v\n", err)
}
