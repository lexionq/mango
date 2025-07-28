package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lexionq/mango/internal/actions"
	"github.com/lexionq/mango/internal/auth"
	"github.com/lexionq/mango/internal/pass"
	"github.com/lexionq/mango/internal/utils"
	"github.com/spf13/cobra"
)

func promptAndVerify(hashPath string) []byte {
	hash, err := auth.GetPassHash(hashPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pass, err := auth.PromptPassword()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ok, err := auth.VerifyPassword(hash, pass)
	if !ok || err != nil {
		fmt.Println("[!] Password incorrect.")
		os.Exit(1)
	}
	return pass
}

func main() {
	homeDir, _ := os.UserHomeDir()
	mangoDir := filepath.Join(homeDir, ".mango")
	mPassPath := filepath.Join(mangoDir, "pass.hash")
	passesPath := filepath.Join(mangoDir, "passes.mango")

	os.MkdirAll(mangoDir, 0700)
	utils.PassesFileControl(passesPath)

	if !utils.PathExists(mPassPath) {
		pass.CreatePasswordandHashPassword(mPassPath)
	}

	var rootCmd = &cobra.Command{
		Use:   "mango",
		Short: "Password manager, written in Golang.🥭",
		Long:  "Password manager, written in Golang.🥭",
	}

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a register",
		Run: func(cmd *cobra.Command, args []string) {
			pass := promptAndVerify(mPassPath)
			actions.AddRegister(passesPath, pass)
		},
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all registers",
		Run: func(cmd *cobra.Command, args []string) {
			pass := promptAndVerify(mPassPath)
			fmt.Println(actions.ListRegisters(passesPath, pass))
		},
	}

	searchCmd := &cobra.Command{
		Use:   "search",
		Short: "Search for a register",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			pass := promptAndVerify(mPassPath)
			word := args[0]
			fmt.Println(actions.Search(passesPath, word, pass))
		},
	}

	editCmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit a register",
		Run: func(cmd *cobra.Command, args []string) {
			pass := promptAndVerify(mPassPath)
			actions.Edit(passesPath, pass)
		},
	}

	importCmd := &cobra.Command{
		Use:   "import",
		Short: "Import registers from a file",
		Run: func(cmd *cobra.Command, args []string) {
			actions.Import(passesPath)
		},
	}

	exportCmd := &cobra.Command{
		Use:   "export",
		Short: "Export all registers to a file",
		Run: func(cmd *cobra.Command, args []string) {
			actions.Export(passesPath)
		},
	}

	rootCmd.AddCommand(addCmd, listCmd, searchCmd, editCmd, importCmd, exportCmd)
	rootCmd.Execute()
}
