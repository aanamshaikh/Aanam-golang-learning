package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type TreeConfig struct {
	rootCmd, dirName string
	dirOnly          bool
	showPermission   bool
	showRelPath      bool
	level            int
	showXML          bool
	showJSON         bool
	sortByTime       bool
}

// type args struct {
// 	cfg                  TreeConfig
// 	info                 fs.FileInfo
// 	permission, fileName string
// 	lenPathSep           int
// }

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
	OpenTag         = "<"
	Slash           = "/"
	CloseTag        = ">"
	Command         = "tree"
	PathSeperator   = string(os.PathSeparator)
	NewLine         = "\n"
)

func main() {
	fmt.Print(tree(input()))
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
	var xml bool
	var json bool
	var sort bool

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
		if x[i] == "-x" {
			xml = true
		}
		if x[i] == "-j" {
			json = true
		}
		if x[i] == "-t" {
			sort = true
		}
	}
	cfg := TreeConfig{rootCmd, dirName, dirOnly, showPermission, showRelPath, level, xml, json, sort}
	return cfg
}

func tree(cfg TreeConfig) string {
	var dir int
	var files int
	var s string

	err := filepath.Walk(cfg.dirName,
		func(path string, info os.FileInfo, err error) error {

			if err != nil {
				return err
			}
			var temp = ""

			fileName := getFileFormatName(cfg.showRelPath, info.Name(), path)
			permission := getPermission(cfg.showPermission, info.Mode().String())
			lenPathSep := len(strings.Split(path, "/"))

			if info.IsDir() {
				dir++
			}
			files++
			if cfg.showXML {
				s = RecInXML(cfg.dirName, temp, 0, cfg, info, permission, fileName)
			} else if cfg.showJSON {
				s = RecInJSON(cfg.dirName, temp, 0, cfg, info, permission, fileName, lenPathSep)
			} else {
				if info.Name() == cfg.dirName {
					dir := fmt.Sprintf("  %v\n", info.Name())
					s += dir
				} else {
					x := getTree(cfg, info, permission, fileName, lenPathSep)
					s += strings.Join(x, ",")
				}
			}
			return nil
		})

	if err != nil {
		log.Fatal(err)
	}

	fileDir := getFilesDir(cfg, files, dir)
	s += fileDir + "\n"
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

func getTreeForLevel(cfg TreeConfig, permission, fileName string, lenPathSep int) string {
	var s string
	if cfg.level > 0 {
		if lenPathSep-1 <= cfg.level {
			s = fmt.Sprintf("%v%v%v%v\n", strings.Repeat(tab,lenPathSep), BoxUpAndRig+BoxHor, permission,fileName)
		}
	} else {
		s = fmt.Sprintf("%v%v%v%v\n", strings.Repeat(tab,lenPathSep), BoxUpAndRig+BoxHor,permission,fileName)
	}
	return s
}

func getFilesDir(cfg TreeConfig, files, dir int) string {
	var s string

	if cfg.dirOnly {
		if cfg.showXML {
			s += "<report>" + NewLine
			s += fmt.Sprintf("%v<directories> %v </directories>%v", strings.Repeat(tab, 3), dir-1, NewLine)
			s += "</report"
		} else if cfg.showJSON {
			s += fmt.Sprintf(",%v%v{type:report,directories:%v}%v]", NewLine, strings.Repeat(tab, 3), dir-1, NewLine)
		} else {
			s = fmt.Sprintf("%v directories ", dir-1)
		}

	} else {
		if cfg.showXML {
			s += "<report>" + NewLine
			s += fmt.Sprintf("%v<directories> %v </directories>%v", strings.Repeat(tab, 3), dir-1, NewLine)
			s += fmt.Sprintf("%v<files> %v </files>%v", strings.Repeat(tab, 3), files, NewLine)
			s += "</report"
		} else if cfg.showJSON {
			s += fmt.Sprintf(",%v%v{type:report,directories:%v,files:%v}%v]", NewLine, strings.Repeat(tab, 3), dir-1, files, NewLine)

		} else {
			s = fmt.Sprintf("%v directories ,%v files\n", dir-1, files)
		}
	}
	return s
}

func getTree(cfg TreeConfig, info fs.FileInfo, permission, fileName string, lenPathSep int) []string {
	var s []string
	if cfg.dirOnly {
		isDir := info.IsDir()
		if isDir {
			x := getTreeForLevel(cfg,permission,fileName,lenPathSep)
			s = append(s, x)
		}
	} else {
		x := getTreeForLevel(cfg,permission,fileName,lenPathSep)
		s = append(s, x)
	}
	return s
}

func RecInJSON(root string, line string, n int, config TreeConfig, info fs.FileInfo, permission, fileName string, lenPathSep int) string {

	files := getFiles(root, config)

	if n == 0 {
		line += leftBrace + NewLine
		line += strings.Repeat(tab, n+2) + " { type: directory ,name:" + root + " " + getPermissions(config, info) + ",contents: " + leftBrace + NewLine
	}

	if n > 0 && n == config.level {
		return line + strings.Repeat(tab, n+2) + "}]" + "\n"
	}

	for _, f := range files {
		if !f.IsDir() {
			line += strings.Repeat(tab, n+5) + "{ type: file ,name:" + f.Name() + " " + getPermissions(config, info) + "}" + NewLine
			continue
		}

		line += strings.Repeat(tab, n+5) + "{ type: directory ,name:" + f.Name() + " " + getPermissions(config, info) + ",contents: [" + NewLine
		fileName := root + "/" + f.Name()
		line = RecInJSON(fileName, line, n+1, config, info, permission, fileName, lenPathSep)
	}

	if n > 0 {
		return line + strings.Repeat(tab, n+4) + "}]" + "\n"
	}

	return line + strings.Repeat(tab, n) + "]" + "\n"
}

func getPermissions(cfg TreeConfig, info fs.FileInfo) string {
	per := ""
	if cfg.showPermission {
		if cfg.showJSON{
			p := info.Mode().String()
			octal := fmt.Sprintf("%#o", info.Mode().Perm())
			per = ",mod :" + octal + ",prot: " + p
		}
		if cfg.showXML{
			p := info.Mode().String()
			octal := fmt.Sprintf("%#o", info.Mode().Perm())
			per = ",mod =" + octal + ",prot= " + p
		}
	}
	return per
}

func getFiles(root string, config TreeConfig) []fs.DirEntry {
	files := ReadDir(root)

	if config.dirOnly {
		files = ReadOnlyDir(files)
	}
	if config.sortByTime {
		SortByModTime(files)
	}
	return files
}

func SortByModTime(files []fs.DirEntry) {
	sort.Slice(files, func(i, j int) bool {
		return getFileInfo(files[i]).ModTime().Unix() < getFileInfo(files[j]).ModTime().Unix()
	})
}

func getFileInfo(fs fs.DirEntry) fs.FileInfo {
	fi, err := fs.Info()
	if err != nil {
		fmt.Println(rightBrace + err.Error() + leftBrace)
		return fi
	}
	return fi
}

func ReadDir(root string) []fs.DirEntry {
	files, err := os.ReadDir(root)
	if err != nil {
		fmt.Println(err)
		return make([]fs.DirEntry, 0)
	}
	return files
}

func ReadOnlyDir(files []fs.DirEntry) []fs.DirEntry {
	dirs := make([]fs.DirEntry, 0)
	for _, f := range files {
		if f.IsDir() {
			dirs = append(dirs, f)
		}
	}
	return dirs
}

func RecInXML(root string, line string, n int, config TreeConfig, info fs.FileInfo, permission, fileName string) string {
	files := getFiles(root, config)

	if n == 0 {
		line += "<tree>" + NewLine
		line += strings.Repeat(tab, n+2) + OpenTag + "directory name= " + root + getPermissions(config, info) + CloseTag + NewLine
	}

	closeDirTag := OpenTag + Slash + "directory" + CloseTag + NewLine
	if n > 0 && n == config.level {
		return line + strings.Repeat(tab, n+3) + closeDirTag
	}

	for _, f := range files {
		if !f.IsDir() {
			line += strings.Repeat(tab, n+4) + OpenTag + "file name="+f.Name()+getPermissions(config, info) + CloseTag +
				OpenTag + Slash + "file" + CloseTag + NewLine
			continue
		}
		line += strings.Repeat(tab, n+4) + OpenTag + "directory name="+f.Name() + getPermissions(config, info) + CloseTag + NewLine
		line = RecInXML(root+"/"+f.Name(), line, n+1, config, info, permission, fileName)
	}
	
	if n > 0 {
		return line + strings.Repeat(tab, n+2) + closeDirTag
	}

	return line + strings.Repeat(tab, n+2) + closeDirTag
}
