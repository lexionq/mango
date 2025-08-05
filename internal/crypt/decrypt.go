package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"github.com/lexionq/mango/internal/auth"
)

func Decrypt(data, key []byte) []byte {
	salt := data[:16]
	key = auth.DeriveKey(key, salt)

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Decrypt error (NewCipher):", err)
		return nil
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("Decrypt error (NewGCM):", err)
		return nil
	}
	nonceSize := aesGCM.NonceSize()

	if len(data) < 16+nonceSize {
		if len(data) == 0 {
			return nil
		} else {
			fmt.Println("Decrypt error: ciphertext(data) too short (len =", len(data), ", needed =", nonceSize, ")")
			return nil
		}
	}

	nonce := data[16 : 16+nonceSize]
	ciphertext := data[16+nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println("Decrypt error (Open):", err)
		return nil
	}

	fmt.Println("[✔] Data decrypted successfully!")

	return plaintext
}
