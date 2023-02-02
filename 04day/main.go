package main

import (
	"flag"
	"fmt"
	"go-program/04day/api"
	"go-program/04day/storage"
	"log"
)

func main() {
	listenAddr := flag.String("listenAddr", ":9090", "teh server address")
	flag.Parse()

	store := storage.NewMemoryStorage()
	server := api.NewServer(*listenAddr, store)
	fmt.Println("Server running on port:", *listenAddr)
	log.Fatalln(server.Start())
}
