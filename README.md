# SIWE Auth (Sign-In With Ethereum for Go)

A lightweight Go module for authenticating users through MetaMask using cryptographic signatures.

## Install

go get github.com/mugabus/siwe-auth

 ## version
 v1.0.0

## Usage
```go
nonce := siwe.GenerateNonce()
ok, err := siwe.VerifySignature(message, signature, address)
