# SIWE Auth (Sign-In With Ethereum for Go)

A lightweight Go module for authenticating users through MetaMask using cryptographic signatures.

## Install

go get github.com/mugabus/siwe-auth


## Usage
```go
nonce := siwe.GenerateNonce()
ok, err := siwe.VerifySignature(message, signature, address)
