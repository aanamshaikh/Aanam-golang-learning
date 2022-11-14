package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type naming struct {
	command  string
	filepath string
}

func getInput() naming {

	fmt.Println("Enter command")
	var command string
	var filepath string
	fmt.Scan(&command, &filepath)

	data := naming{command, filepath}
	return data
}

func searchRecursively(file string) {
	err := filepath.Walk(file,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// fmt.Println(info.Name())
			rel, err := filepath.Rel(file, path)
			fmt.Println(strings.Repeat(rel,1))
			
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
}
func main() {

	data := getInput()
	// fmt.Println("command: ", data.command, "\nfilepath:", data.filepath)
searchRecursively(data.filepath)

	
	

}
