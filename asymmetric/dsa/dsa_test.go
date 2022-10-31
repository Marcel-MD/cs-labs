package dsa

import (
	"crypto/rand"
	"math/big"
	"testing"
)

func TestValidDsaSignVerify(t *testing.T) {
	params := new(Parameters)
	if err := GenerateParameters(params, rand.Reader); err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}

	priv := new(PrivateKey)
	priv.Parameters = *params
	err := GenerateKey(priv, rand.Reader)
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}

	msg := []byte("learning")
	r, s, err := Sign(rand.Reader, priv, msg)
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}

	if !Verify(&priv.PublicKey, msg, r, s) {
		t.Fatalf("Expected signature to be valid but got invalid")
	}
}

func TestInvalidDsaSignVerify(t *testing.T) {
	params := new(Parameters)
	if err := GenerateParameters(params, rand.Reader); err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}

	priv := new(PrivateKey)
	priv.Parameters = *params
	err := GenerateKey(priv, rand.Reader)
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}

	msg := []byte("learning")
	r, s, err := Sign(rand.Reader, priv, msg)
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}

	// Change the signature
	s.Add(s, big.NewInt(1))

	if Verify(&priv.PublicKey, msg, r, s) {
		t.Fatalf("Expected signature to be invalid but got valid")
	}
}
