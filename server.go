package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Server!")
	})
	// const serverAddr string = "127.0.0.1:3000"
	const port string = ":3000"

	fmt.Println("Server Listening on port 3000")
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting server", err)
	}
}
