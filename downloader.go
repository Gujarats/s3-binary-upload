package main

import (
	"fmt"
	"os"
	"os/exec"
)

func Downloader(links []string, username, password string) {

	app := "wget"
	user := "--user"
	userValue := username
	passwordOpt := "--password"
	passwordValue := password
	recursive := "-r"
	exclude := "-R"
	excludeValue := "index.html*"
	noParentDir := "-np"
	noOverwrite := "-nc"
	noHost := "-nH"

	for _, link := range links {
		fmt.Println("downloading all artifacts from = ", link)
		runCommand(app, user, userValue, passwordOpt, passwordValue, recursive, exclude, excludeValue, noParentDir, noOverwrite, noHost, link)
	}

	fmt.Println("all download success")
}

func runCommand(cmdName string, arg ...string) []byte {
	cmd := exec.Command(cmdName, arg...)

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to start command %s. %s\n", cmdName, err.Error())
		os.Exit(1)
	}

	return []byte(``)

}
