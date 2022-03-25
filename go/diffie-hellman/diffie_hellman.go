package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

func PrivateKey(p *big.Int) *big.Int {
    n, _ := rand.Int(rand.Reader, new(big.Int).Sub(p, big.NewInt(2)))
    return n.Add(n, big.NewInt(2))
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
    return new(big.Int).Exp(big.NewInt(g), private, p)
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
    private := PrivateKey(p)
	return private, PublicKey(private, p, g)
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
