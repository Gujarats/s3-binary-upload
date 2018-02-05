package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"

	"github.com/kr/fs"
)

const (
	slash = "/"
	dot   = "."
)

//filter all directory to specific package name
func filterDir(source []byte, filter string) []string {
	var result []string
	directories := strings.Split(string(source), "\n")
	for _, dir := range directories {
		if strings.Contains(string(dir), filter) {
			result = append(result, string(dir))
		}
	}
	return result
}

func getFilesPathFrom(path string) []string {
	var filesPath []string
	walker := fs.Walk(path)
	for walker.Step() {
		if err := walker.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if !walker.Stat().IsDir() {
			filesPath = append(filesPath, walker.Path())
		}
	}

	return filesPath
}

func homeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

func getGradleCacheDir() string {
	return homeDir() + gradleCacheDir
}
