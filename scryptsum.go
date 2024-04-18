package main

import (
	"flag"
	"fmt"
	"golang.org/x/crypto/scrypt"
)

func main() {
	var L, N, r, p int

	flag.IntVar(&L, "L", 32, "")

	flag.IntVar(&N, "N", 21, "")
	flag.IntVar(&r, "r", 8, "")
	flag.IntVar(&p, "p", 1, "")

	flag.Parse()

	var password string
	fmt.Scanln(&password)

	fmt.Println(" settings:\n N r p / L\n", N, r, p, "/", L)

	hash, err := scrypt.Key([]byte(password), []byte{}, 1<<uint(N), r, p, L)
	if err != nil {
		fmt.Printf("error generating hash: %v\n", err)
		return
	}

	fmt.Printf("%x\n", hash)
}
