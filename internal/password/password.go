package password

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lexionq/mango/internal/auth"
	"github.com/lexionq/mango/internal/crypt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

func CreatePasswordandHashPassword(path string){
	fmt.Print("Create password: ")
	pass, err0 := term.ReadPassword(int(os.Stdin.Fd()))
	if err0 != nil {
		fmt.Println(err0)
	}
	fmt.Println(" ")

	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	
	if err != nil {
		fmt.Println(err)
	}

	file, err1 := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)

	if err1 != nil {
		fmt.Println(err1)
	}

	defer file.Close()
	
	_, err2 := file.Write(hash)

	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println("[✔] Password created and saved successfully!")


}

func ChangePassword() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	mangoDir := filepath.Join(homeDir,".mango")
	const mPassFile = "pass.hash"
	const passesFile = "passes.mango"
	mPassPath := filepath.Join(mangoDir,mPassFile)
	passesPath := filepath.Join(mangoDir,passesFile)

	fmt.Print("Enter password again: ")
	pass, err0 := term.ReadPassword(int(os.Stdin.Fd()))
	if err0 != nil {
		fmt.Println(err0)
	}
	fmt.Println(" ")

	hash,err := auth.GetPassHash(mPassPath)
	if err != nil {
		fmt.Println(err)
	}

	okay, err := auth.VerifyPassword(hash,pass)
	if okay {
		fmt.Println("[✔] Password true.")
	} else {
		fmt.Println("[!] Password false.")
		os.Exit(1)
	}

	if err != nil{
		fmt.Println("[!] Error: ", err)
	}

	fmt.Print("Enter your new pass: ")
	newPass, err1 := term.ReadPassword(int(os.Stdin.Fd()))
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(" ")

	newPassHash, err := bcrypt.GenerateFromPassword(newPass, bcrypt.DefaultCost)
	
	if err != nil {
		panic(err)
	}

	file, err1 := os.OpenFile(mPassPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)

	if err1 != nil {
		panic(err1)
	}

	defer file.Close()
	
	_, err2 := file.Write(newPassHash)

	if err2 != nil {
		panic(err2)
	}

	fmt.Println("[✔] New password saved successfully. ")

	fmt.Print("Enter your new pass again for verification: ")
	passs, err1 := term.ReadPassword(int(os.Stdin.Fd()))
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(" ")

	okay, err = auth.VerifyPassword(newPassHash,passs)
	if okay {
		fmt.Println("[✔] Password true.")
	} else {
		fmt.Println("[!] Password false.")
		os.Exit(1)
	}
	if err != nil {
		fmt.Println("[!] Error: ",err)
	}

	ciphertext,err := os.ReadFile(passesPath)
	if err != nil  {
		fmt.Println("[!] Error: ",err)
	}

	plaintext := crypt.Decrypt(ciphertext,pass)

	newCiphertext := crypt.Encrypt(plaintext,passs)

	f, err := os.OpenFile(passesPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("File couldn't open.:", err)
		return
	}
	defer f.Close()

	_, err = f.Write(newCiphertext)

	if err != nil {
		fmt.Println("Error: ",err)
	}

	fmt.Println("[✔] Your password updated successfully!")
	

}