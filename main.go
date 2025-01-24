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

	fmt.Println("Original Message:", string(message))

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

	// Decode the hexadecimal HMAC back to bytes
	decodedHMAC, err := hex.DecodeString(hmacHex)
	if err != nil {
		fmt.Println("Error decoding HMAC:", err)
		return
	}

	fmt.Println("Decoded Original Message:", string(message))

	verifyHMAC := hmac.New(sha256.New, key)
	verifyHMAC.Write(message)
	computedHMAC := verifyHMAC.Sum(nil)
	if hmac.Equal(decodedHMAC, computedHMAC) {
		fmt.Println("HMAC verification successful: Decoded HMAC matches computed HMAC")
	} else {
		fmt.Println("HMAC verification failed: Decoded HMAC does not match computed HMAC")
	}

}
