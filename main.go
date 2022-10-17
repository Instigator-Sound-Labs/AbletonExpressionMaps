package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	convert "github.com/basgys/goxml2json"
)

func main() {

}

func getDir() {
	// get current Dir from the walk.
	path, _ := os.Getwd()
	fmt.Println(path)
}

func walkPath() { // maybe move to main?
	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("cannot get current dir: %v\n", err)
		return
	}
	os.Chdir(path)

	subDirToSkip := "skip"

	fmt.Println("On Unix:")
	// ensure the files are only plist files.
	err = filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() && info.Name() == subDirToSkip {
			fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
			return filepath.SkipDir
		}
		fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", path, err)
		return
	}
}

func makeJson() {
	// xml is an io.Reader
	xml := strings.NewReader(`<?xml version="1.0" encoding="UTF-8"?><hello>world</hello>`)
	json, err := convert.Convert(xml)
	if err != nil {
		panic("That's embarrassing...")
	}

	fmt.Println(json.String())
	// {"hello": "world"}
}

func convertJson() {
	// json to json
}

func makeFile() {
	// make file for each file traveresd
}
