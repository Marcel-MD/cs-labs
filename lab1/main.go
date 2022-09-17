package main

import (
	"fmt"
)

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	shift := 3
	key := "secret"
	var seed int64 = 10
	text := "hello world"

	encrypted := caesarPermutationEncrypt(alphabet, seed, shift, text)
	fmt.Printf("Caesar encrypted: %s\n", encrypted)
	decrypted := caesarPermutationDecrypt(alphabet, seed, shift, encrypted)
	fmt.Printf("Caesar decrypted: %s\n", decrypted)

	encrypted = vigenereEncrypt(alphabet, key, text)
	fmt.Printf("Vigenere encrypted: %s\n", encrypted)
	decrypted = vigenereDecrypt(alphabet, key, encrypted)
	fmt.Printf("Vigenere decrypted: %s\n", decrypted)

	encrypted = playfairEncrypt(key, text)
	fmt.Printf("Playfair encrypted: %s\n", encrypted)
	decrypted = playfairDecrypt(key, encrypted)
	fmt.Printf("Playfair decrypted: %s\n", decrypted)
}
