package main

import (
	"encoding/json"
	"net/http"

	"github.com/mugabus/siwe-auth/siwe"
)

var currentNonce string

func main() {

	// Route: Generate nonce
	http.HandleFunc("/api/nonce", func(w http.ResponseWriter, r *http.Request) {
		currentNonce = siwe.GenerateNonce()
		json.NewEncoder(w).Encode(map[string]string{"nonce": currentNonce})
	})

	// Route: Verify signature
	http.HandleFunc("/api/verify", func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			Message   string `json:"message"`
			Signature string `json:"signature"`
			Address   string `json:"address"`
		}
		json.NewDecoder(r.Body).Decode(&body)

		ok, err := siwe.VerifySignature(body.Message, body.Signature, body.Address)
		if err != nil || !ok {
			http.Error(w, "invalid signature", 401)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	http.ListenAndServe(":8080", nil)
}
