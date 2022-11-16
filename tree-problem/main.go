package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	//Box Drawing Characters
	BoxVer       = "│"
	BoxHor       = "──"
	BoxVH        = BoxVer + BoxHor
	BoxDowAndRig = "┌"
	BoxDowAndLef = "┐"
	BoxUpAndRig  = "└"
	BoxUpAndLef  = "┘"
)

type Command struct {
	command string
	// option
	// level
	filepath string
}

type options struct {

}

func getInput() Command {

	fmt.Println("Enter command")
	var command string
	var filepath string
	fmt.Scan(&command, &filepath)

	data := Command{command, filepath}
	return data
}

func searchRecursively(file string) {
	var totalDir int
	var totalfiles int
	err := filepath.Walk(file,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			x := len(strings.Split(path, "/"))

			if info.Name() == file {
				fmt.Printf("  %v\n", info.Name())
			} else {
				fmt.Printf("%v%v%v\n", strings.Repeat("  ", x), BoxUpAndRig+BoxHor, info.Name())
			}

			if info.IsDir() {
				totalDir++
			}
			totalfiles++
			// fmt.Print("|")
			return nil
		})
	fmt.Printf("%v directories ,%v files\n", totalDir-1, totalfiles)

	if err != nil {
		log.Fatal(err)
	}
}

func searchRecursivelyWithFilePermisssion(file string) {
	var totalDir int
	var totalfiles int
	err := filepath.Walk(file,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			rel, err := filepath.Rel(file, path)
			x := len(strings.Split(rel, "/"))

			if info.IsDir() {
				totalDir++
			}
			totalfiles++
			// fmt.Printf("%v%v\n",strings.Repeat("  ",x),info.Name())

			if info.Name() == file {
				fmt.Printf("%v", strings.Repeat("                 ", x))
				fmt.Printf("%v\n", info.Name())
			} else {
				// calculate ln of permission add with x
				fmt.Printf("%v%v", strings.Repeat("                 ", x), BoxUpAndRig+BoxHor)
				fmt.Printf(" [%v] %v\n", info.Mode(), info.Name())
			}

			return nil
		})
	fmt.Printf("%v directories ,%v files\n", totalDir-1, totalfiles)

	if err != nil {
		log.Fatal(err)
	}
}

func searchDirOnly(file string) {
	var totalDir int

	err := filepath.Walk(file,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			x := len(strings.Split(path, "/"))
			if info.Name() == file {
				fmt.Printf("%v", strings.Repeat("   ", x))
				fmt.Printf("%v\n", info.Name())
			} else {
				if info.IsDir() {
					totalDir++
					fmt.Printf("%v%v%v\n", strings.Repeat("  ", x), BoxUpAndRig+BoxHor, info.Name())

				}
			}

			return nil
		})
	fmt.Printf("%v directories\n", totalDir-1)

	if err != nil {
		log.Fatal(err)
	}
}

func searchRecursivelyWithRelativePath(file string) {
	var totalDir int
	var totalfiles int
	err := filepath.Walk(file,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			x := len(strings.Split(path, "/"))

			if info.Name() == file {

				fmt.Printf("  %v\n", info.Name())
			} else {

				fmt.Printf("%v%v%v\n", strings.Repeat("  ", x), BoxUpAndRig+BoxHor, path)

			}

			if info.IsDir() {
				totalDir++
			}
			totalfiles++
			// fmt.Print("|")
			return nil
		})
	fmt.Printf("%v directories ,%v files\n", totalDir-1, totalfiles)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	data := getInput()

	fmt.Println("#----------Recursive Search-----------#")
	searchRecursively(data.filepath)

	// fmt.Println("#----------Recursive Search For Dir -----------#")
	// searchDirOnly(data.filepath)

	// fmt.Println("#----------Recursive Search with file permission-----------#")
	// searchRecursivelyWithFilePermisssion(data.filepath)

	// fmt.Println("#----------Recursive Search with relative path-----------#")
	// searchRecursivelyWithRelativePath(data.filepath)

}
