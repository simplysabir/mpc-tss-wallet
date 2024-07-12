package crypto

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

type Signer interface {
    Sign(message []byte) ([]byte, error)
}

type Secp256k1Signer struct {
    privateKey *ecdsa.PrivateKey
}

func (s *Secp256k1Signer) Sign(message []byte) ([]byte, error) {
    // In a real implementation, you would use the TSS library to perform the signing
    // This is a simplified example using a single private key
    if s.privateKey == nil {
        var err error
        s.privateKey, err = crypto.GenerateKey()
        if err != nil {
            return nil, fmt.Errorf("error generating secp256k1 key: %v", err)
        }
    }

    signature, err := crypto.Sign(crypto.Keccak256(message), s.privateKey)
    if err != nil {
        return nil, fmt.Errorf("error signing with secp256k1: %v", err)
    }

    return signature, nil
}

type Ed25519Signer struct {
    privateKey ed25519.PrivateKey
}

func (s *Ed25519Signer) Sign(message []byte) ([]byte, error) {
    // In a real implementation, you would use the TSS library to perform the signing
    // This is a simplified example using a single private key
    if s.privateKey == nil {
        var err error
        _, s.privateKey, err = ed25519.GenerateKey(nil)
        if err != nil {
            return nil, fmt.Errorf("error generating ed25519 key: %v", err)
        }
    }

    signature := ed25519.Sign(s.privateKey, message)
    return signature, nil
}