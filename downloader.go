package main

import (
	"fmt"
	"os"
	"os/exec"
)

func Downloader(links []string, username, password string) {

	// full command
	//wget -A jar,pom,xml,md5,sha1 -m -p -E -k -K -np -nH --user YOUR_USER --password YOUR_PASS YOUR_URL
	app := "wget"

	accept := "-A"
	acceptValues := "jar,pom,xml,md5,sha1"

	// authentication
	user := "--user"
	userValue := username
	passwordOpt := "--password"
	passwordValue := password

	//required options
	mirror := "-m"
	pageRequisite := "-p"
	adjustExtensions := "-E"
	convertLinks := "-k"
	backupConverted := "-K"
	//recursive := "-r"
	//exclude := "-R"
	//excludeValue := "index.html*"
	noParentDir := "-np"
	noHost := "-nH"

	for _, link := range links {
		fmt.Println("downloading all artifacts from = ", link)
		runCommand(app, accept, acceptValues, user, userValue, passwordOpt, passwordValue, mirror, pageRequisite, adjustExtensions, convertLinks, backupConverted, noParentDir, noHost, link)
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
