package vigenere

import "testing"

func TestEncryptDecrypt(t *testing.T) {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	key := "secret"
	text := "hello world"

	encrypted := Encrypt(alphabet, key, text)
	if string(text) == string(encrypted) {
		t.Fatalf("Expected encrypted text to be different from %s", text)
	}

	decrypted := Decrypt(alphabet, key, encrypted)
	if string(text) != string(decrypted) {
		t.Fatalf("Expected decrypted text to be %s text but got %s", text, decrypted)
	}
}
