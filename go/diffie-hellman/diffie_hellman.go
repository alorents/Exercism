package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// PrivateKey generates a random key with value greater than 1 and less than p
func PrivateKey(p *big.Int) *big.Int {
	_one := big.NewInt(1)
	privateKey := _one
	for privateKey.Cmp(_one) <= 0 {
		var err error
		privateKey, err = rand.Int(rand.Reader, p)
		if err != nil {
			return big.NewInt(0)
		}
	}

	return privateKey
}

// PublicKey calulates a public key = g**a mod p
func PublicKey(a, p *big.Int, g int64) *big.Int {
	_one := big.NewInt(1)
	publicKey := _one
	publicKey.Exp(big.NewInt(g), a, p)
	return publicKey
}

// SecretKey calculates a secret key = B**a mod p
func SecretKey(a, B, p *big.Int) *big.Int {
	_one := big.NewInt(1)
	secretKey := _one
	secretKey.Exp(B, a, p)

	return secretKey
}

// NewPair generates a new private key and public key from provided p, and g values
func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	a := PrivateKey(p)
	A := PublicKey(a, p, g)
	return a, A
}
