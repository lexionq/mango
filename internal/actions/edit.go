package actions

import (
	"fmt"
	"github.com/lexionq/mango/internal/crypt"
	"os"
	"os/exec"
	"strings"
)


func Edit(path string, key []byte) error {
	ciphertext,err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("[!] failed to read file: %w", err)
	}

	plaintext := crypt.Decrypt(ciphertext, key)
	fmt.Println("Registers:",string(plaintext))
	
	tmpFile, err := os.CreateTemp("","edit-*.txt")
	if err != nil{
		panic(err)
	}
	defer os.Remove(tmpFile.Name())

	initialText := string(plaintext)
	_, err = tmpFile.WriteString(initialText)
	if err != nil{
		panic(err)
	}
	tmpFile.Close()

	cmd := exec.Command("nano",tmpFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	err = cmd.Run()
	if err != nil {
		fmt.Println("[!] If you are using Windows, Mac or you don't have `nano`, you can't use this feature!")
		panic(err)
	}

	data,err := os.ReadFile(tmpFile.Name())
	if err != nil {
		panic(err)
	}	

	datas := strings.Split(string(data),",")
	
	if len(datas) != 5{
		fmt.Println("[!] Error: The minimum number of entries must be at least 4 and at most 5.")
		os.Exit(1)
	}

	text := crypt.Encrypt(data,key)

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("File couldn't open.:", err)
		return err
	}
	defer f.Close()

	_, err = f.Write(text)

	if err != nil {
		fmt.Println("Error: ",err)
	}

	fmt.Println("[✔] Registers edited successfully!")

	fmt.Println("[*] New data:\n",string(data))
	return nil
}