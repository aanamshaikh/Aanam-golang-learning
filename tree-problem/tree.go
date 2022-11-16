package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type options struct {
	rootCmd, dirName string
	option           option

	// dir string
	// permission string
	// level, filename string
	// level_no int
}

type option struct {
	name, desc string
	value      int
}

const (
	//Box Drawing Characters
	BoxVer       = "│"
	BoxHor       = "──"
	BoxVH        = BoxVer + BoxHor
	BoxDowAndRig = "┌"
	BoxDowAndLef = "┐"
	BoxUpAndRig  = "└"
	BoxUpAndLef  = "┘"
	leftBrace    = "["
	rightBrace   = "]"
)

func getFileFormatName(fileNameFormat, name, path string) string {
	fileName := name
	if fileNameFormat == "-f" {
		fileName = path
	}
	return fileName
}

func getPermission(permission, mode string) string {
	per := ""
	if permission == "-p" {
		m := mode
		per = leftBrace + m + rightBrace
	}
	return per
}


func tree(dirName string, level, permission, dir, fileNameFormat option) {
	// var totalDir int
	// var totalfiles int

	err := filepath.Walk(dirName,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			fileName := getFileFormatName(fileNameFormat.name, info.Name(), path)

			permission := getPermission(permission.name, info.Mode().String())

			x := len(strings.Split(path, "/"))

			if info.Name() == dirName {

				fmt.Printf("  %v\n", info.Name())

			} else {
				// extract to func
				if dir.name == "-d" {
					if info.IsDir() {
						if level.value != 0 {
							if x-1 <= level.value {
								fmt.Printf("%v%v%v%v\n", strings.Repeat("  ", x), BoxUpAndRig+BoxHor, permission, fileName)
							}
						}else{
							fmt.Printf("%v%v%v%v\n", strings.Repeat("  ", x), BoxUpAndRig+BoxHor, permission, fileName)
					    }
				    }
				} else {
					if level.value != 0 {
						if x-1 <= level.value {
							fmt.Printf("%v%v%v%v\n", strings.Repeat("  ", x), BoxUpAndRig+BoxHor, permission, fileName)
						}
				    }else{
					fmt.Printf("%v%v%v%v\n", strings.Repeat("  ", x), BoxUpAndRig+BoxHor, permission, fileName)
				    }
			}
		}
		
			return nil

		})

	// 	if(dir.name=="-d"){
	// 	    fmt.Printf("%v directories ,%v files\n",)
	// 	}
	// fmt.Printf("%v directories ,%v files\n", totalDir-1, totalfiles)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	var level = option{"-l", "l", 1}
	permission := option{"-p", "p", 0}
	dir := option{"-d", "d", 0}
	fileNameFormat := option{"-f", "f", 0}

	if level.value<1 {
		fmt.Println(" level should be greater than 0")
		
	}
	tree("dir", level, permission, dir, fileNameFormat)

}
