package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"log"

	"github.com/btcsuite/btcutil/base58"
)

func main() {
	// Generate private key.
	key, _, _, err := elliptic.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("%s\n", base58.Encode(key))
}
