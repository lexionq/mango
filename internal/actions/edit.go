package actions

import (
	"bufio"
	"fmt"
	"github.com/lexionq/mango/internal/crypt"
	"os"
	"os/exec"
	"strings"
)

func Edit(path string, key []byte) error {
	ciphertext, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("[!] failed to read file: %w", err)
	}

	plaintext := crypt.Decrypt(ciphertext, key)
	fmt.Println("Registers:", string(plaintext))

	tmpFile, err := os.CreateTemp("", "edit-*.txt")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpFile.Name())

	initialText := string(plaintext)
	_, err = tmpFile.WriteString(initialText)
	if err != nil {
		panic(err)
	}
	tmpFile.Close()

	// Cross-platform editor selection
	var cmd *exec.Cmd
	editor := os.Getenv("EDITOR")

	if editor != "" {
		cmd = exec.Command(editor, tmpFile.Name())
	} else {
		// Try different editors based on OS and availability
		editors := []string{"nano", "vim", "vi", "notepad", "code", "gedit"}

		for _, ed := range editors {
			if _, err := exec.LookPath(ed); err == nil {
				cmd = exec.Command(ed, tmpFile.Name())
				break
			}
		}

		if cmd == nil {
			fmt.Println("[!] No suitable editor found. Please set EDITOR environment variable.")
			fmt.Println("    Available editors: nano, vim, vi, notepad, code, gedit")
			return fmt.Errorf("no editor available")
		}
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Printf("[!] Error running editor: %v\n", err)
		return err
	}

	editedData, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(editedData)))
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		fields := strings.Split(scanner.Text(), ",")
		if len(fields) < 4 || len(fields) > 5 {
			fmt.Println("[!] Error: The minimum number of entries must be at least 4 and at most 5.")
			return fmt.Errorf("invalid number of fields in line %d", lineNum)
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error scanning edited data: %w", err)
	}

	text := crypt.Encrypt(editedData, key)

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("File couldn't open.:", err)
		return err
	}
	defer f.Close()

	_, err = f.Write(text)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("[✔] Registers edited successfully!")

	fmt.Println("[*] New data:\n", string(editedData))
	return nil
}
