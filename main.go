package main

import (
	"fmt"
	"log"
	"axolot/src/encryption"
)

func main() {
	// Call the GenerateKeyFromHostDetails function to generate the key
	key, err := encryption.GenerateKeyFromHostDetails()
	if err != nil {
		log.Fatalf("Error generating key: %v", err)
	}

	// Print the generated key (in raw byte form)
	fmt.Printf("Generated key: %x\n", key) // Using %x to print the byte slice as a hex string

	// Success message
	fmt.Println("Key has been successfully generated")
}
