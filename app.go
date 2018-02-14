package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"

	"github.com/Gujarats/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	S3_REGION      = "ap-southeast-1"
	S3_BUCKET      = "s3-website-test.hashicorp.com"
	gradleCacheDir = "/.gradle/caches/modules-2/files-2.1"
)

// TODO : put this in config file
var S3_BUCKETS = []string{
	"gujarats-test1",
	"gujarats-test2",
	"gujarats-test3",
}

func main() {
	// Create a single AWS session (we can re use this if we're uploading many files)
	s, err := session.NewSession(&aws.Config{Region: aws.String(S3_REGION)})
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

	// TODO : create a CLI app using cobra
	// - create config file to specify s3 buckets single or multiple
	// - speficy gradle cache location
	// - upload binrary artifact to different s3 random or sequencial

	// TODO : fast way
	// - get all s3 and upload the files separately

	// list all the directory names
	command := exec.Command("ls")
	command.Dir = path.Join(homeDir(), gradleCacheDir)
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
	for _, artifact := range artifacts {
		logger.Debug("artifact :: ", artifact)
		s3Bucket := nextS3(S3_BUCKETS)
		for _, fileDir := range artifact {
			err = AddFileToS3(s, fileDir, s3Bucket)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

// AddFileToS3 will upload a single file to S3, it will require a pre-built aws session
// and will set file info like content type and encryption on the uploaded file.
func AddFileToS3(s *session.Session, fileDir string, s3Bucket string) error {

	// Open the file for use
	file, err := os.Open(fileDir)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get file size and read the file content into a buffer
	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	// Modify the fileDirectory to custom dir so it can be downloaded by gradle
	removedEncDir := removeEncryptPath(fileDir)
	newFileDir := folderBuilder(removedEncDir)
	logger.Debug("newFileDir :: ", newFileDir)

	// Config settings: this is where you choose the bucket, filename, content-type etc.
	// of the file you're uploading.

	_, err = s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(s3Bucket),
		Key:                  aws.String(newFileDir),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})

	return err
}
