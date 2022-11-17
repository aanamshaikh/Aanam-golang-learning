package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type TreeConfig struct {
	rootCmd, dirName string
	dirOnly          bool
	showPermission   bool
	showRelPath      bool
	level            int
}

type args struct {
	cfg                  TreeConfig
	info                 fs.FileInfo
	permission, fileName string
	lenPathSep           int
}

const (
	//Box Drawing Characters
	BoxVer          = "│"
	BoxHor          = "──"
	BoxVH           = BoxVer + BoxHor
	BoxDowAndRig    = "┌"
	BoxDowAndLef    = "┐"
	BoxUpAndRig     = "└"
	BoxUpAndLef     = "┘"
	leftBrace       = "["
	rightBrace      = "]"
	tab             = "  "
	dirOnlyC        = "-d"
	showPermissionC = "-p"
	showRelPathC    = "-f"
	levelC          = "-l"
)

func main() {
	
	printTree(tree(input()))
}

func input() TreeConfig {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	treeConfig := parseCommand(line)
	return treeConfig
}
func parseCommand(line string) TreeConfig {

	x := strings.Split(line, " ")
	rootCmd := x[0]
	dirName := x[len(x)-1]
	var dirOnly bool
	var showPermission bool
	var showRelPath bool
	level := 0

	for i := 1; i < len(x)-1; i++ {

		if x[i] == dirOnlyC {
			dirOnly = true
		}
		if x[i] == showPermissionC {
			showPermission = true
		}
		if x[i] == showRelPathC {
			showRelPath = true
		}
		if x[i] == levelC {
			level = parseToInt(x[i+1])
		}
	}
	cfg := TreeConfig{rootCmd, dirName, dirOnly, showPermission, showRelPath, level}
	return cfg
}

func tree(cfg TreeConfig) []string {
	var dir int
	var files int
	var s []string

	err := filepath.Walk(cfg.dirName,
		func(path string, info os.FileInfo, err error) error {

			if err != nil {
				return err
			}

			fileName := getFileFormatName(cfg.showRelPath, info.Name(), path)
			permission := getPermission(cfg.showPermission, info.Mode().String())
			lenPathSep := len(strings.Split(path, "/"))
			args := args{cfg, info, permission, fileName, lenPathSep}

			if info.IsDir() {
				dir++
			}
			files++

			if info.Name() == cfg.dirName {
				dir := fmt.Sprintf("  %v\n", info.Name())
				s = append(s, dir)
			} else {
				s = append(s, getTree(args)...)
			}
			return nil
		})

	if err != nil {
		log.Fatal(err)
	}

	fileDir := getFilesDir(cfg, files, dir)
	s = append(s, fileDir+"\n")
	return s
}

func parseToInt(input string) int {
	num, err := strconv.ParseInt(input, 10, 32)
	//check for error
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return int(num)
}

func getFileFormatName(fileNameFormat bool, name, path string) string {
	fileName := name
	if fileNameFormat {
		fileName = path
	}
	return fileName
}

func getPermission(permission bool, mode string) string {
	per := ""
	if permission {
		m := mode
		per = leftBrace + m + rightBrace
	}
	return per
}

func getTreeForLevel(args args) string {
	var s string
	if args.cfg.level > 0 {
		if args.lenPathSep-1 <= args.cfg.level {
			s = fmt.Sprintf("%v%v%v%v\n", strings.Repeat(tab, args.lenPathSep), BoxUpAndRig+BoxHor, args.permission, args.fileName)
		}
	} else {
		s = fmt.Sprintf("%v%v%v%v\n", strings.Repeat(tab, args.lenPathSep), BoxUpAndRig+BoxHor, args.permission, args.fileName)
	}
	return s
}

func getFilesDir(cfg TreeConfig, files, dir int) string {
	var s string
	if cfg.dirOnly {
		s = fmt.Sprintf("%v directories ", dir-1)
	} else {
		s = fmt.Sprintf("%v directories ,%v files\n", dir-1, files)
	}
	return s
}

func getTree(args args) []string {
	var s []string
	if args.cfg.dirOnly {
		isDir := args.info.IsDir()
		if isDir {
			x := getTreeForLevel(args)
			s = append(s, x)
		}
	} else {
		x := getTreeForLevel(args)
		s = append(s, x)
	}
	return s
}

func printTree(tree []string) {
	for _, t := range tree {
		fmt.Print(t)
	}
}
