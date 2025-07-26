package utils

import (
	"os"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func PassesFileControl(path string){
	if !PathExists(path){
		os.Create(path)
	}
}

