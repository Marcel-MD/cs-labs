package rsa

import (
	"crypto/rand"
	"io"
	"math/big"
)

// PublicKey represents an RSA public key.
type PublicKey struct {
	N *big.Int // modulus
	E *big.Int // public exponent
}

// PrivateKey represents an RSA private key.
type PrivateKey struct {
	PublicKey
	D *big.Int // private exponent
}

// GenerateKey generates an RSA keypair of the given bit size using the
// random source rand (for example, crypto/rand.Reader).
func GenerateKey(r io.Reader, bits int) (priv *PrivateKey, err error) {

	p, err := rand.Prime(r, bits/2)
	if err != nil {
		return
	}

	q, err := rand.Prime(r, bits/2)
	if err != nil {
		return
	}

	n := new(big.Int).Mul(p, q)                                                                     // n = p * q
	phi := new(big.Int).Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1))) // phi = (p-1) * (q-1)
	e := big.NewInt(3)                                                                              // e = 3
	d := new(big.Int).ModInverse(e, phi)                                                            // d = e^-1 mod phi

	priv = &PrivateKey{PublicKey{N: n, E: e}, d}

	return
}

func Encrypt(pub *PublicKey, msg []byte) []byte {
	m := new(big.Int).SetBytes(msg)
	c := new(big.Int).Exp(m, pub.E, pub.N)
	return c.Bytes()
}

func Decrypt(priv *PrivateKey, c []byte) []byte {
	ci := new(big.Int).SetBytes(c)
	m := new(big.Int).Exp(ci, priv.D, priv.N)
	return m.Bytes()
}
