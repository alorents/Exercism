package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// PrivateKey generates a random key with value greater than 1 and less than p
func PrivateKey(p *big.Int) *big.Int {
	two := big.NewInt(2)
	max := new(big.Int).Sub(p, two)
	privateKey, err := rand.Int(rand.Reader, max)
	if err != nil {
		return nil
	}
	return privateKey.Add(privateKey, two)
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
