package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
)

func main() {
	pb := pocketbase.NewWithConfig(&pocketbase.Config{})

	if err := pb.Start(); err != nil {
		log.Fatal(err)
	}
}
