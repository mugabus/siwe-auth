SIWE Auth (Sign-In With Ethereum for Go)

A lightweight and production-ready Go module for authenticating users with MetaMask using EIP-191 / EIP-4361 signatures.

This library provides:

✅ Secure nonce generation
✅ Signature verification using go-ethereum
✅ Works with any Web3-enabled frontend
✅ Small, dependency-light, and easy to integrate

Perfect for modern API backends, dApps, dashboards, and Web3 login systems.

🚀 Installation
go get github.com/mugabus/siwe-auth

📌 Version

Latest stable release:

v1.0.0

🔧 Usage Example
Generate a Nonce
nonce := siwe.GenerateNonce()


Send this nonce to the frontend, where the user signs it with MetaMask.

Verify a Signature
ok, err := siwe.VerifySignature(message, signature, address)
if err != nil {
    log.Println("Invalid signature:", err)
}

if ok {
    fmt.Println("✓ Signature verified, user authenticated!")
}

🧠 How It Works

Backend generates a random nonce

Frontend asks MetaMask to sign a message containing that nonce

Backend runs VerifySignature() to:

Recover the signer’s public key

Convert it to an Ethereum address

Compare it with the expected address

If they match → User is authenticated

This avoids passwords entirely and uses cryptographic proof-of-ownership.

📁 Library Structure
/siwe-auth
   ├── nonce.go      → Generate secure nonces
   ├── verify.go     → Recover signer address + verify
   ├── utils.go      → Helpers (optional)
   ├── README.md
   └── LICENSE

🔒 Security Notes

Always store & validate nonces server-side

Use HTTPS in production

Nonces should be single-use and expire quickly

Do not accept signatures for unsent messages

🧩 Frontend Example (MetaMask)
const accounts = await window.ethereum.request({
  method: "eth_requestAccounts",
});

const from = accounts[0];

const signature = await ethereum.request({
  method: "personal_sign",
  params: [message, from],
});


Send message, address, and signature to your Go backend.

🏆 Why This Library?

No heavy frameworks

No blockchain node required

Pure cryptographic Ethereum signature validation

Production-ready for APIs

Lightweight & minimal — just what backend devs need

📄 License

This project is released under the MIT License.
See LICENSE for full details.

❤️ Contributing

Pull requests and issues are welcome!
Let’s improve Web3 authentication for everyone.

⭐ Support the Project

If you like this library, give it a GitHub star 🌟
It helps boost visibility for other developers.