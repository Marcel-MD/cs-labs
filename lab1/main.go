package main

import (
	"fmt"

	"github.com/Marcel-MD/cs-labs/caesar"
	"github.com/Marcel-MD/cs-labs/playfair"
	"github.com/Marcel-MD/cs-labs/vigenere"
)

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	shift := 3
	key := "secret"
	var seed int64 = 10
	text := "hello world"

	encrypted := caesar.PermutationEncrypt(alphabet, seed, shift, text)
	fmt.Printf("Caesar encrypted: %s\n", encrypted)
	decrypted := caesar.PermutationDecrypt(alphabet, seed, shift, encrypted)
	fmt.Printf("Caesar decrypted: %s\n", decrypted)

	encrypted = vigenere.Encrypt(alphabet, key, text)
	fmt.Printf("Vigenere encrypted: %s\n", encrypted)
	decrypted = vigenere.Decrypt(alphabet, key, encrypted)
	fmt.Printf("Vigenere decrypted: %s\n", decrypted)

	encrypted = playfair.Encrypt(key, text)
	fmt.Printf("Playfair encrypted: %s\n", encrypted)
	decrypted = playfair.Decrypt(key, encrypted)
	fmt.Printf("Playfair decrypted: %s\n", decrypted)
}
