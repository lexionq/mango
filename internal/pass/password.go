package pass

import (
	"fmt"
	"os"
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