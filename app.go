package main

import (
	"fmt"
	"log"
	"os/exec"
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

// TODO ; upload the file wihtous cache
// user can choose which artifact to upload in gradle cache
// or specifically choose the path
func main() {
	config := getConfig()
	// Create a single AWS session (we can re use this if we're uploading many files)
	s, err := session.NewSession(&aws.Config{Region: aws.String(config.Region)})
	if err != nil {
		log.Fatal(err)
	}

	// get package name
	fmt.Print("\nEnter your package = ")
	var packageName string
	_, err = fmt.Scan(&packageName)
	if err != nil {
		log.Fatal(err)
	}

	// list all the directory names
	command := exec.Command("ls")
	command.Dir = path.Join(homeDir(), config.GradleCacheDir)
	packageNames, err := command.Output()
	if err != nil {
		log.Fatal(err)
	}

	// store arifact jar & pom
	// key = artifact name prefix without extenstion like com.traveloka.common/accessor-1.0.2
	// []string all files dir
	artifacts := make(map[string][]string)

	// get specific directory for scanning artifact
	packages := filterDir(packageNames, packageName)
	for _, pack := range packages {
		files := getFilesPathFrom(pack)
		for _, file := range files {
			artifactName := getArtifactName(file)
			// store artifact
			artifacts[artifactName] = append(artifacts[artifactName], file)
		}
	}

	// Upload
	var buckets []string
	buckets = append(buckets, config.S3Bucket)
	buckets = append(buckets, config.S3Buckets...)
	upload(s, buckets, artifacts)
}
