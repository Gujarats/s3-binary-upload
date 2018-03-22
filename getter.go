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
// if given filter equals to empty string then all package will be return
func filterDir(source []byte, filter string) []string {
	var result []string
	directories := strings.Split(string(source), "\n")
	if filter == "all" {
		return directories
	}

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

//TODO : move path of gradle caches to somewhere else
// because the path could change
func getGradleCacheDir() string {
	return homeDir() + "/.gradle/caches/modules-2/files-2.1"
}

func getArtifactName(fileDir string) string {
	parts := strings.Split(fileDir, ".")
	result := strings.Join(parts[:len(parts)-1], ".")
	return result
}

func getArtifactNameForGradle(fileDir string) string {
	result := getArtifactName(fileDir)
	result = removeEncryptPath(result)
	return result
}
