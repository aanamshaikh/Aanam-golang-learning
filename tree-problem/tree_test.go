package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrintDirOnly(t *testing.T) {

	want := []string{
		tab + "dir\n",
		tab + tab + BoxUpAndRig + BoxHor + "subdir1\n",
		tab + tab + tab + BoxUpAndRig + BoxHor + "subdir\n",
		tab + tab + BoxUpAndRig + BoxHor + "subdir2\n",
	}

	cfg := TreeConfig{"tree", "dir", true, false, false, 0}
	got := tree(cfg)
	assert := assert.New(t)
	for i := 0; i < len(want); i++ {
		assert.Equal(want[i], got[i], "got and want are not equal")

	}
}

func TestPrintWithPermissions(t *testing.T) {

	want := []string{
		tab + "dir\n",
		tab + tab + BoxUpAndRig + BoxHor + "[-rw-rw-r--]file23\n",
		tab + tab + BoxUpAndRig + BoxHor + "[drwxrwxr-x]subdir1\n",
		tab + tab + tab + BoxUpAndRig + BoxHor + "[drwxrwxr-x]subdir\n",
	}

	cfg := TreeConfig{"tree", "dir", false, true, false, 0}
	got := tree(cfg)

	assert := assert.New(t)
	for i := 0; i < len(want); i++ {
		assert.Equal(want[i], got[i], "got and want are not equal")
	}
}

func TestPrintWithRelPath(t *testing.T) {

	want := []string{
		tab + "dir\n",
		tab + tab + BoxUpAndRig + BoxHor + "dir/file23\n",
		tab + tab + BoxUpAndRig + BoxHor + "dir/subdir1\n",
		tab + tab + tab + BoxUpAndRig + BoxHor + "dir/subdir1/subdir\n",
	}

	cfg := TreeConfig{"tree", "dir", false, false, true, 0}
	got := tree(cfg)

	assert := assert.New(t)
	for i := 0; i < len(want); i++ {
		assert.Equal(want[i], got[i], "got and want are not equal")

	}
}

func TestPrintWithLevel(t *testing.T) {

	want := []string{
		tab + "dir\n",
		tab + tab + BoxUpAndRig + BoxHor + "file23\n",
		tab + tab + BoxUpAndRig + BoxHor + "subdir1\n",
	}

	cfg := TreeConfig{"tree", "dir", false, false, false, 1}
	got := tree(cfg)

	assert := assert.New(t)
	for i := 0; i < len(want); i++ {
		assert.Equal(want[i], got[i], "got and want are not equal")
	}
}

func TestPrintWithAllFlags(t *testing.T) {

	want := []string{
		tab + "dir\n",
		tab + tab + BoxUpAndRig + BoxHor + leftBrace + "drwxrwxr-x" + rightBrace + "dir/subdir1\n",
		tab + tab + BoxUpAndRig + BoxHor + "[drwxrwxr-x]dir/subdir2\n",
	}

	cfg := TreeConfig{"tree", "dir", true, true, true, 1}
	got := tree(cfg)

	assert := assert.New(t)
	for i := 0; i < len(want); i++ {
		assert.Equal(want[1], got[1], "got and want are not equal")

	}

}
