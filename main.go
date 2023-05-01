package main

import (
	"log"
)

func main() {
	store, err := NewPostGressStore()
	if err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":3000", store)
	server.Run()
}
