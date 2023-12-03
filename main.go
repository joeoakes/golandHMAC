package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	k := "Joes Key"
	fmt.Println("Keyword: " + k)
	s := "{'payload':'hashLab'}"
	fmt.Println("Message sent to the receiver: " + s)
	p := k + s
	fmt.Println("Key+Payload: " + p)
	h := sha256.New()
	h.Write([]byte(p))
	bs := h.Sum(nil)
	fmt.Printf("Hash value sent to the receiver: %x\n", bs)

}
