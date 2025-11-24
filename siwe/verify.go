package siwe

import (
	"encoding/hex"
	"errors"

	"github.com/ethereum/go-ethereum/crypto"
)

func VerifySignature(message, signatureHex, expectedAddress string) (bool, error) {
	sig, err := hex.DecodeString(signatureHex[2:])
	if err != nil {
		return false, err
	}

	// Normalize the recovery ID (last byte)
	if sig[64] >= 27 {
		sig[64] -= 27
	}

	hash := crypto.Keccak256Hash([]byte(message))

	pubKey, err := crypto.SigToPub(hash.Bytes(), sig)
	if err != nil {
		return false, err
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey).Hex()

	if recoveredAddr != expectedAddress {
		return false, errors.New("signature does not match address")
	}

	return true, nil
}
