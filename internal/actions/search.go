package actions

import (
	"fmt"
	"github.com/lexionq/mango/internal/crypt"
	"os"
	"strings"
)

func Clean(s string) string {
	return strings.TrimSpace(strings.ReplaceAll(s, "\r", ""))
}



func Search(path,word string, key []byte) []string{
	if word == ""{
		fmt.Println("[!] Error: Search key too short!")
		return nil
	}

	ciphertext,err := os.ReadFile(path)
	if err != nil {
		fmt.Println("[!] Error: ",err)
	}
	var result []string

	plaintext := crypt.Decrypt(ciphertext,key)
	if plaintext == nil {
		fmt.Println("[!] Error: Decryption failed or returned nil.")
		return nil
	}

	text := strings.Split(string(plaintext),"\n")

	for i, line := range text {
		if Clean(line) == ""{
			continue
		}
		wordss := strings.Split(line,",")
		for _, wordInLine := range wordss {
			if strings.Contains(Clean(wordInLine), Clean(word)) {
				result = append(result, fmt.Sprintf("Found in Line %d: %s", i+1, line))
			break
			}
		}

	}
	return result
}