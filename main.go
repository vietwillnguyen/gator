package main

import (
	"fmt"
	"gator/internal/config"
	"log"
)

func main() {
	fmt.Println("Gator Application Start.")
	config, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	fmt.Printf("Before Write: Config: %v\n", config)
	config.SetUser("Viet")
	fmt.Printf("After Write: Config: %v\n", config)
}
