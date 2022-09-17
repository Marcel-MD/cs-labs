package main

import (
	"fmt"
)

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	shift := 3
	var seed int64 = 10
	text := "hello world"

	encrypted := caesarPermutationEncrypt(alphabet, seed, shift, text)
	decrypted := caesarPermutationDecrypt(alphabet, seed, shift, encrypted)

	fmt.Println(encrypted)
	fmt.Println(decrypted)
}
