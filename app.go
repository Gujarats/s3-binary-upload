package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

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

	// list all the directory names
	// TODO : ls command at specific directory not current directory
	packageNames, err := exec.Command("ls").Output()
	if err != nil {
		log.Fatal(err)
	}

	// get specific directory for scanning artifact
	packages := filterDir(packageNames, packageName)
	for _, pack := range packages {
		files := getFilesPathFrom(pack)
		for _, file := range files {
			// Upload
			err = AddFileToS3(s, file)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}

// AddFileToS3 will upload a single file to S3, it will require a pre-built aws session
// and will set file info like content type and encryption on the uploaded file.
func AddFileToS3(s *session.Session, fileDir string) error {

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
	logger.Debug("fileDir = ", fileDir)
	logger.Debug("newFileDir = ", newFileDir)

	// Config settings: this is where you choose the bucket, filename, content-type etc.
	// of the file you're uploading.
	_, err = s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(S3_BUCKET),
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
