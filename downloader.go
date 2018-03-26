package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
)

func Downloader(links []string, username, pass string) {

	// full command
	//wget -A jar,pom,xml,md5,sha1 -m -p -E -k -K -np -nH --user YOUR_USER --password YOUR_PASS YOUR_URL
	app := "wget"

	accept := "-A"
	acceptValues := "jar,pom,xml,md5,sha1"

	// authentication
	user := "--user"
	userValue := username
	password := "--password"
	passwordValue := pass

	//required options
	mirror := "-m"
	pageRequisite := "-p"
	adjustExtensions := "-E"
	convertLinks := "-k"
	backupConverted := "-K"
	noParentDir := "-np"
	noHost := "-nH"
	directoryPrefix := "-P"
	directoryPrefixValue := path.Join(getHomeDir(), ".s3-binary-upload")

	for _, link := range links {
		fmt.Println("downloading all artifacts from = ", link)
		runCommand(app, "", false, accept, acceptValues, user, userValue, password, passwordValue, directoryPrefix, directoryPrefixValue, mirror, pageRequisite, adjustExtensions, convertLinks, backupConverted, noParentDir, noHost, link)
	}

	fmt.Println("all download success")
}

// parameter print will determine we output the result or not
func runCommand(cmdName string, dir string, print bool, arg ...string) []byte {
	cmd := exec.Command(cmdName, arg...)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if dir != "" {
		cmd.Dir = dir
	}

	if print {
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
	}

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to start command %s. %s\n", cmdName, err.Error())
		os.Exit(1)
	}

	return stdout.Bytes()

}
