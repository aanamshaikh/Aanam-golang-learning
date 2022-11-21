package main

import (
	"testing"
)

func TestPrintDirOnly(t *testing.T) {

	want := []string{
		"dir",
		tab+BoxDowAndRig+BoxHor+"subdir1",
		tab+tab+BoxDowAndRig+BoxHor+"subdir",
		tab+BoxDowAndRig+BoxHor+"subdir2",
	}

	cfg := TreeConfig{"tree","dir",true,false,false,0}
	got:=tree(cfg)

	for i:=0; i<len(want); i++ {
		if got[i] != want[i] {
			t.Errorf("got %v does not match  want %v ",got[i],want[i])
		}
	}
}

func TestPrintWithPermissions(t *testing.T) {

	want := []string{
		"dir",
		tab+BoxDowAndRig+BoxHor+"[-rw-rw-r--]"+"file23",
		tab+BoxDowAndRig+BoxHor+"[dwxrwxr-x]"+"subdir1",
		tab+tab+BoxDowAndRig+BoxHor+"[dwxrwxr-x]"+"subdir",
	}

	cfg := TreeConfig{"tree","dir",false,true,false,0}
	got:=tree(cfg)

	for i:=0; i<len(want); i++ {
		if got[i] != want[i] {
			t.Errorf("got %v does not match  want %v ",got[i],want[i])
		}
	}
}

func TestPrintWithRelPath(t *testing.T) {

	want := []string{
		"dir",
		tab+BoxDowAndRig+BoxHor+"dir/file23",
		tab+BoxDowAndRig+BoxHor+"dir/subdir1",
		tab+tab+BoxDowAndRig+BoxHor+"dir/subdir1/subdir",
	}

	cfg := TreeConfig{"tree","dir",false,false,true,0}
	got:=tree(cfg)

	for i:=0; i<len(want); i++ {
		if got[i] != want[i] {
			t.Errorf("got %v does not match  want %v ",got[i],want[i])
		}
	}
}

func TestPrintWithLevel(t *testing.T) {

	want := []string{
		"dir",
		tab+BoxDowAndRig+BoxHor+"file23",
		tab+BoxDowAndRig+BoxHor+"subdir1",
		tab+BoxDowAndRig+BoxHor+"subdir2",
	}

	cfg := TreeConfig{"tree","dir",false,false,false,1}
	got:=tree(cfg)

	for i:=0; i<len(want); i++ {
		if got[i] != want[i] {
			t.Errorf("got %v does not match  want %v ",got[i],want[i])
		}
	}
}

func TestPrintWithAllFlags(t *testing.T) {

	want := []string{
		"dir",
		tab+BoxDowAndRig+BoxHor+"[dwxrwxr-x]"+"dir/subdir1",
		tab+BoxDowAndRig+BoxHor+"[dwxrwxr-x]"+"dir/subdir2",
	}

	cfg := TreeConfig{"tree","dir",true,true,true,1}
	got:=tree(cfg)

	for i:=0; i<len(want); i++ {
		if got[i] != want[i] {
			t.Errorf("got %v does not match  want %v ",got[i],want[i])
		}
	}
}

