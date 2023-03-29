package main

import (
	"fmt"
	"net/http"

	"github.com/bhusan-shumsher/go-web-training/pkg/handlers"
)

const port = ":8000"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf(fmt.Sprintf("Starting application on port %s", port))

	_ = http.ListenAndServe(port, nil)
}
