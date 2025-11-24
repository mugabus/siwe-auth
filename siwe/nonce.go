package siwe

import "github.com/google/uuid"

func GenerateNonce() string {
	return uuid.New().String()
}
