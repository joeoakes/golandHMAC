package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	k := "keyword"
	s := "{'payload':'hashLab'}"
	p := k + s
	h := sha256.New()
	h.Write([]byte(p))
	bs := h.Sum(nil)
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
