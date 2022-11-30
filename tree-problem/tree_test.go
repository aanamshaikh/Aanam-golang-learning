package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrintDirOnly(t *testing.T) {
	want := tab + "dir\n" + tab + tab + BoxUpAndRig + BoxHor + "subdir1\n" + tab + tab + tab + BoxUpAndRig + BoxHor + "subdir\n" + tab + tab + BoxUpAndRig + BoxHor + "subdir2\n" + "3 directories \n"

	cfg := TreeConfig{"tree", "dir", true, false, false, 0, false, false, false}
	got := tree(cfg)
	assert := assert.New(t)
	assert.Equal(want, got)
}

func TestPrintWithPermissions(t *testing.T) {
	want := "  dir\n    └──[-rw-rw-r--]file23\n    └──[drwxrwxr-x]subdir1\n      └──[drwxrwxr-x]subdir\n        └──[-rw-rw-r--]vectoe\n      └──[-rw-rw-r--]subfile1\n    └──[drwxrwxr-x]subdir2\n      └──[-rw-rw-r--]file2\n3 directories ,8 files\n\n"
	cfg := TreeConfig{"tree", "dir", false, true, false, 0, false, false, false}
	got := tree(cfg)

	assert := assert.New(t)
	assert.Equal(want, got)
}

func TestPrintWithRelPath(t *testing.T) {
	want := "  dir\n    └──dir/file23\n    └──dir/subdir1\n      └──dir/subdir1/subdir\n        └──dir/subdir1/subdir/vectoe\n      └──dir/subdir1/subfile1\n    └──dir/subdir2\n      └──dir/subdir2/file2\n3 directories ,8 files\n\n"
	cfg := TreeConfig{"tree", "dir", false, false, true, 0, false, false, false}
	got := tree(cfg)

	assert := assert.New(t)
	assert.Equal(want, got)
}

func TestPrintWithLevel(t *testing.T) {
	want := "  dir\n    └──file23\n    └──subdir1\n    └──subdir2\n3 directories ,8 files\n\n"
	cfg := TreeConfig{"tree", "dir", false, false, false, 1, false, false, false}
	got := tree(cfg)

	assert := assert.New(t)
	assert.Equal(want, got)
}

func TestPrintWithAllFlags(t *testing.T) {
	want := "  dir\n    └──[drwxrwxr-x]dir/subdir1\n    └──[drwxrwxr-x]dir/subdir2\n3 directories \n"
	cfg := TreeConfig{"tree", "dir", true, true, true, 1, false, false, false}
	got := tree(cfg)

	assert := assert.New(t)
	assert.Equal(want, got)
}

func TestPrintJsonTree(t *testing.T) {
	want := "[\n     { type: directory ,name:dir ,contents: [\n          { type: directory ,name:subdir1 ,contents: [\n            { type: directory ,name:subdir ,contents: [\n            }]\n          }]\n          { type: directory ,name:subdir2 ,contents: [\n          }]\n]\n,\n      {type:report,directories:3}\n]\n"
	cfg := TreeConfig{"tree", "dir", true, false, false, 0, false, true, false}
	got := tree(cfg)

	assert := assert.New(t)
	assert.Equal(want, got)
}

func TestPrintXMLTree(t *testing.T) {
    want:="<tree>\n    <directory name= dir>\n        <directory name=subdir1>\n          <directory name=subdir>\n        </directory>\n      </directory>\n        <directory name=subdir2>\n      </directory>\n    </directory>\n<report>\n      <directories> 3 </directories>\n</report\n"
	cfg := TreeConfig{"tree", "dir", true, false, false, 0, true, false, false}
	got := tree(cfg)

	assert := assert.New(t)
	assert.Equal(want, got)
}

func TestPrintWithSortTime(t *testing.T) {
	want := "  dir\n    └──subdir1\n      └──subdir\n    └──subdir2\n3 directories \n"
	cfg := TreeConfig{"tree", "dir", true, false, false, 0, false, false, true}
	got := tree(cfg)

	assert := assert.New(t)
	assert.Equal(want, got)
}
