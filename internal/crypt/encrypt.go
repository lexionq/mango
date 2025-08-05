package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"github.com/lexionq/mango/internal/auth"
	"io"
)

func Encrypt(data, key []byte) []byte {
	salt, err0 := auth.GenerateSalt(16)
	if err0 != nil {
		fmt.Println("Encrypt error (GenerateSalt): ", err0)
	}
	if _, errr := io.ReadFull(rand.Reader, salt); errr != nil {
		fmt.Println("Encrypt error (GenerateSalt):", errr)
		return nil
	}

	key = auth.DeriveKey(key, salt)

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Encrypt error (NewCipher):", err)
		return nil
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("Encrypt error (NewGCM):", err)
		return nil
	}
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println("Encrypt error (ReadFull):", err)
		return nil
	}
	ciphertext := aesGCM.Seal(nil, nonce, data, nil)

	result := append(salt, nonce...)
	result = append(result, ciphertext...)

	fmt.Println("[✔] Data encrypted successfully.")

	return result
}
