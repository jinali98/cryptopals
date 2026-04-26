package c03singlebytexorcipher

import (
	"cryptopals/set_1/c01_hex_to_base64"
	"fmt"
	"log"
)

//heuristic English-likeness scorer
func ScoreEnglish(text string) int {
	score := 0
	for _, c := range text {
		switch {
		case c == ' ':
			score += 3
		case c == 'e' || c == 't' || c == 'a' || c == 'o' || c == 'i' || c == 'n':
			score += 2
		case (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z'):
			score += 1
		}
	}
	return score
}

func SingleByteXOR(input []byte, key byte) []byte {
	result := make([]byte, len(input))
	for i := range input {
		result[i] = input[i] ^ key
	}
	return result
}

func Solution() (string, error) {
	hexString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	bytes, err := c01_hex_to_base64.HexToBytes(hexString)
	if err != nil {
		log.Fatal(err)
	}

	bestScore := 0
	bestResult := ""
	bestKey := byte(0)

	/*

Here's what's in the 0-255 range:

0–31: Control characters (things like newline, tab — not printable)
32–47: Space, punctuation (!, ", #, $, etc.)
48–57: Digits 0 through 9
65–90: Uppercase A through Z
97–122: Lowercase a through z
128–255: Extended ASCII (accented characters, symbols)

	*/

	for i := 0; i < 256; i++ {
		key := byte(i)
		decrypted := SingleByteXOR(bytes, key)
		text := string(decrypted)
		score := ScoreEnglish(text)

		if score > bestScore {
			bestScore = score
			bestResult = text
			bestKey = key
		}
	}

	fmt.Printf("Key: %c (0x%02x)\n", bestKey, bestKey)
	fmt.Printf("Decrypted: %s\n", bestResult)

	return bestResult, nil
}