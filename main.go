package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed version.txt
var version string

//go:embed logo.png
var logo []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := ioutil.WriteFile("logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dirEntries, _ := path.ReadDir("files")
	for _, enrty := range dirEntries {
		if !enrty.IsDir() {
			fmt.Println(enrty.Name())
			content, _ := path.ReadFile("files/" + enrty.Name())
			fmt.Println("Content:", string(content))
		}
	}
}
