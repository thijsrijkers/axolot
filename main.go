package main

import (
	"fmt"
	"log"
	"os"
	"axolot/src/encryption"
)

func main() {
	if len(os.Args) > 1 {
		key := os.Args[1]
		keyBytes := []byte(key)
		encryption.SetKey(keyBytes)

		fmt.Printf("Using provided key. %x\n", "")
	} else {
		key, err := encryption.GenerateKeyFromHostDetails()
		if err != nil {
			log.Fatalf("Error generating key: %v", err)
		}

		fmt.Printf("Generated key: %x\n", key)
	}
}
