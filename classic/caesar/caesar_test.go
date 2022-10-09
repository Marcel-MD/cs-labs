package caesar

import "testing"

func TestEncryptDecrypt(t *testing.T) {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	shift := 3
	text := "hello world"

	encrypted := Encrypt(alphabet, shift, text)
	if string(text) == string(encrypted) {
		t.Fatalf("Expected encrypted text to be different from %s", text)
	}

	decrypted := Decrypt(alphabet, shift, encrypted)
	if string(text) != string(decrypted) {
		t.Fatalf("Expected decrypted text to be %s text but got %s", text, decrypted)
	}
}
