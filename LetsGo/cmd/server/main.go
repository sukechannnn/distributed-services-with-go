package main

import (
	"log"

	"github.com/sukechannnn/distributed-services-with-go/LetsGo/server"
)

func main() {
	log.Printf("Start server...")
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
