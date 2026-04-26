package main

import (
	c03singlebytexorcipher "cryptopals/set_1/c03_single_byte_xor_cipher"
	"log"
)

func main() {
	
	_, err := c03singlebytexorcipher.Solution()
	if err != nil {
		log.Fatal(err)
	}
}