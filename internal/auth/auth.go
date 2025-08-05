package auth

import (
	"crypto/rand"
	"fmt"
	"os"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

func GetPassHash(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func PromptPassword() ([]byte, error) {
	fmt.Print("Enter password: ")
	pass, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	return pass, err
}

func VerifyPassword(hash []byte, pass []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hash, pass)
	if err != nil {
		return false, nil // password is false
	}
	return true, nil // password is true
}

func GenerateSalt(size int) ([]byte, error) {
	salt := make([]byte, size)
	_, err := rand.Read(salt)
	return salt, err
}

func DeriveKey(password, salt []byte) []byte {
	return argon2.IDKey(password, salt, 1, 64*1024, 4, 32)
}
