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
	http.HandleFunc("/getDetails/", http2.GetDetails)
	http.HandleFunc("/getRestDetails/", http2.GetRestDetails)
	http.HandleFunc("/makeComment/", http2.MakeComment)

	var port = ":" + os.Getenv("PORT")
	if port == ":" {
		log.Println("Port not found")
		port = ":8080"
	}

	fmt.Println("Started server at http://localhost" + port)
	log.Fatalln(http.ListenAndServe(port, nil))
}
