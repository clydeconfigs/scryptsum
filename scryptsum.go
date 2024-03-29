package main

import (
	"flag"
	"fmt"
	"golang.org/x/crypto/scrypt"
)

func main() {
	var L, N, r, p int

	flag.IntVar(&L, "L", 64, "")

	flag.IntVar(&N, "N", 20, "")
	flag.IntVar(&r, "r", 8, "")
	flag.IntVar(&p, "p", 4, "")

	flag.Parse()

	var password string
	fmt.Scanln(&password)

	hash, err := scrypt.Key([]byte(password), []byte{}, 1<<uint(N), r, p, L)
	if err != nil {
		fmt.Printf("error generating hash: %v\n", err)
		return
	}

	fmt.Printf("%x\n", hash)
}
