package c01_hex_to_base64

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

/*
Cryptopals Rule
Always operate on raw bytes, never on encoded strings. Only use hex and base64 for pretty-printing.
*/
func Solution() (string, error) {
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"


	// convert the  hex string to bytes
	bytes, err := HexToBytes(hexString)
	if err != nil {
		log.Fatal(err)
	}

	// convert the bytes to base64 string
	base64String := BytesToBase64(bytes)
	fmt.Println(base64String)
	return base64String, nil
}

func HexToBytes(hexString string) ([]byte, error) {
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func BytesToBase64(bytes []byte) string {
	base64String := base64.StdEncoding.EncodeToString(bytes)
	return base64String
}

func BytesToHex(bytes []byte) string {
	hexString := hex.EncodeToString(bytes)
	return hexString
}