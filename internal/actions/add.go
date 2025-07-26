package actions

import (
	"bufio"
	"fmt"
	"mango/internal/crypt"
	"os"
	"strings"
	"syscall"
	"golang.org/x/term"
)

func AddRegister(path string, key []byte ){
	var rgName, rgSite, rgUsername, rgNote string
	var rgPass []byte

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Register Name: ")
	input, _ := reader.ReadString('\n')
	rgName = strings.TrimSpace(input)

	fmt.Print("Website: ")
	input, _ = reader.ReadString('\n')
	rgSite = strings.TrimSpace(input)

	fmt.Print("Username: ")
	input, _ = reader.ReadString('\n')
	rgUsername = strings.TrimSpace(input)

	fmt.Print("Password: ")
	rgPass, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()

	if err != nil {
		fmt.Println("Error reading password:", err)
		return
	}

	fmt.Print("Notes (optional): ")
	input, _ = reader.ReadString('\n')
	rgNote = strings.TrimSpace(input)

	ciphertext,err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error: ",err)
	}	

	var plaintext []byte

	if len(ciphertext) > 0 {
		plaintext = crypt.Decrypt(ciphertext,key)
	} else {
		plaintext = []byte{}
	}

	metaStr := rgName + "," + rgSite + "," + rgUsername + "," 
	text := append(plaintext, []byte(metaStr)...)
	text = append(text, rgPass...)
	text = append(text, []byte(","+rgNote+"\n")...)

	text_to_be_written := crypt.Encrypt(text,key)

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("File couldn't open.:", err)
		return
	}
	defer f.Close()

	_, err = f.Write(text_to_be_written)

	if err != nil {
		fmt.Println("Error: ",err)
	}

	fmt.Println("[✔] Register added successfully!")

	
}

