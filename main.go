package main

import (
	"context"
	"log"
	"os"

	"github.com/go-shop/config"
)

func main() {
	ctx := context.Background()
	_ = ctx

	// Initialize config

	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error .env path required")
		}
		return os.Args[1]
	}())

	log.Println(cfg)

}
