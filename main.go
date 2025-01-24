package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	// Key and message
	key := []byte("cryptii")
	message := []byte("Hello")

	// Create a new HMAC using SHA-256
	h := hmac.New(sha256.New, key)

	// Write the message to the HMAC
	h.Write(message)

	// Compute the HMAC signature
	signature := h.Sum(nil)

	// Convert the HMAC to a hexadecimal string for display
	hmacHex := hex.EncodeToString(signature)

	// Print the HMAC
	fmt.Println("HMAC (SHA-256):", hmacHex)
}
