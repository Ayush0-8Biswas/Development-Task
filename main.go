package main

import (
	http2 "eateries-in-kgp/pkg/http"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./cmd/web/")))
	http.HandleFunc("/getRestaurants/", http2.GetRestaurants)
	fmt.Println("Started server at http://localhost:8080")
	var port = os.Getenv("port")
	if port == "" {
		port = ":8080"
	}
	log.Fatalln(http.ListenAndServe(port, nil))
}
