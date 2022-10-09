package playfair

import "testing"

func TestEncryptDecrypt(t *testing.T) {
	key := "secret"
	text := "cryptography"

	encrypted := Encrypt(key, text)
	if string(text) == string(encrypted) {
		t.Fatalf("Expected encrypted text to be different from %s", text)
	}

	decrypted := Decrypt(key, encrypted)
	if string(text) != string(decrypted) {
		t.Fatalf("Expected decrypted text to be %s text but got %s", text, decrypted)
	}
}
