package actions

import (
	"bufio"
	"fmt"
	cp "github.com/otiai10/copy"
	"os"
	"path/filepath"
	"strings"
)

func Import(dest_path string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter source path for importing: ")
	source_path, _ := reader.ReadString('\n')
	source_path = strings.TrimSpace(source_path)

	err := cp.Copy(source_path, dest_path)
	if err != nil {
		fmt.Println("[!] Error during import: ", err)
	} else {
		fmt.Println("[✔] Importing successfully finished!")
	}
}

func Export(source_path string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter destination directory for exporting: ")
	dest_path, _ := reader.ReadString('\n')
	dest_path = strings.TrimSpace(dest_path)

	destPath := filepath.Join(dest_path, "passes.mango")

	err := cp.Copy(source_path, destPath)
	if err != nil {
		fmt.Println("[!] Error during export: ", err)
	} else {
		fmt.Println("[✔] Exporting successfully finished! ")
	}
}
