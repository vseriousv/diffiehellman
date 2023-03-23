package transport_key

import (
	"encoding/hex"
	"errors"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"math/big"
)

// GetTransportKey computes a shared transport key using ECDH with the secp256k1 curve.
// publicKey: hex-encoded public key of the other party (prefixed with "0x").
// privateKey: hex-encoded private key (prefixed with "0x").
func GetTransportKey(publicKey, privateKey string) (string, error) {
	privateKeyAHex := privateKey[2:]
	privateKeyABigInt, success := new(big.Int).SetString(privateKeyAHex, 16)
	if !success {
		return "", errors.New("invalid private key format")
	}

	publicKeyBHex := publicKey[2:]
	publicKeyBBytes, err := hex.DecodeString(publicKeyBHex)
	if err != nil {
		return "", errors.New("invalid public key format")
	}
	publicKeyB, err := crypto.UnmarshalPubkey(publicKeyBBytes)
	if err != nil {
		return "", errors.New("unmarshalling public key failed")
	}

	sharedSecretX, _ := secp256k1.S256().ScalarMult(publicKeyB.X, publicKeyB.Y, privateKeyABigInt.Bytes())
	if sharedSecretX == nil {
		return "", errors.New("scalar multiplication failed")
	}

	transportKey := crypto.Keccak256(sharedSecretX.Bytes())
	return hex.EncodeToString(transportKey), nil
}
