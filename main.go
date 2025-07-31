package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lexionq/mango/internal/actions"
	"github.com/lexionq/mango/internal/auth"
	"github.com/lexionq/mango/internal/password"
	"github.com/lexionq/mango/internal/utils"
)


func main(){
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	mangoDir := filepath.Join(homeDir,".mango")
	const mPassFile = "pass.hash"
	const passesFile = "passes.mango"


	if !utils.PathExists(mangoDir){
		os.Mkdir(mangoDir, 0700)
	}

	mPassPath := filepath.Join(mangoDir,mPassFile)
	passesPath := filepath.Join(mangoDir,passesFile)

	if !utils.PathExists(mPassPath){
		password.CreatePasswordandHashPassword(mPassPath)
	}

	utils.PassesFileControl(passesPath)
	

	hash,err := auth.GetPassHash(mPassPath)
	if err != nil {
		fmt.Println(err)
	}

	var action string
	if len(os.Args) > 1 {
		action = os.Args[1]
	} else {
		utils.DisplayHelpMessage()
		return
	}

	pass,err := auth.PromptPassword()

	okay, err1 := auth.VerifyPassword(hash, pass)

	if err1 != nil {
		fmt.Println(err)
	}													

	if okay {
		fmt.Println("[✔] Password true.")
	} else {
		fmt.Println("[!] Password false. | ", err1)
		os.Exit(1)
	}	

	

	switch action {
	case "add":
		actions.AddRegister(passesPath,pass)
	case "list":
		fmt.Println("\n",actions.ListRegisters(passesPath,pass))
	case "search":
		word := os.Args[2]
		fmt.Println(actions.Search(passesPath,word,pass))
	case "edit":
		actions.Edit(passesPath,pass)
	case "export":
		actions.Export(passesPath)
	case "change":
		password.ChangePassword()
	case "help","-h","--help":
		utils.DisplayHelpMessage()	
	default:
		fmt.Println("[!] Error Unknown parameter!")
		utils.DisplayHelpMessage()
	}

}