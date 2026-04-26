package c02fixedxor

import (
	"cryptopals/set_1/c01_hex_to_base64"
	"errors"
	"fmt"
)


func Solution( string1 string, string2 string ) (string, error) {
	bytes1, err := c01_hex_to_base64.HexToBytes(string1)
	if err != nil {
		return "", err
	}
	bytes2, err := c01_hex_to_base64.HexToBytes(string2)
	if err != nil {
		return "", err
	}
	// if the bytes are not the same length, return an error
	if len(bytes1) != len(bytes2) {
		return "", errors.New("bytes are not the same length")
	}

	// the new xor bytes slice. the length is the same as the original bytes
	xorBytes := make([]byte, len(bytes1))

	// loop through the bytes and xor them
	for i := range bytes1 {
		xorBytes[i] = bytes1[i] ^ bytes2[i]
	}

	// convert the xor bytes to a hex string
	xorString := c01_hex_to_base64.BytesToHex(xorBytes)
	fmt.Println(xorString)
	return xorString, nil
}

