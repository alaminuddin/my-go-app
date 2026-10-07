package main

import (
	"fmt"
	"net/http"
)

func main() {

	server := http.NewServeMux()

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	fmt.Println("Server is running on http://localhost:8086...")

	if err := http.ListenAndServe(":8086", server); err != nil {
		panic(err)
	}

}
