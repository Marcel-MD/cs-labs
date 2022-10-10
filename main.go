package main

import (
	"fmt"

	"github.com/Marcel-MD/cs-labs/stream-block/rc4"
)

func main() {
	fmt.Println("Hello World!")

	key := []byte("secret")
	text := []byte("hello world")

	encrypted := make([]byte, len(text))
	rc4.XORKeyStream(encrypted, text, key)
	fmt.Printf("RC4 Encrypted: %s\n", encrypted)

	decrypted := make([]byte, len(encrypted))
	rc4.XORKeyStream(decrypted, encrypted, key)
	fmt.Printf("RC4 Decrypted: %s\n", decrypted)
}
