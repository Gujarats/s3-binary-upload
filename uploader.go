package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/Gujarats/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// pass buckets to uploader to start uploading
func upload(s *session.Session, config *Config, buckets []string, artifacts map[string][]string, isGradleDir bool) {
	for _, artifact := range artifacts {
		s3Bucket := nextS3(buckets)
		for _, fileDir := range artifact {
			buffer, contentLength := getFileSize(fileDir)

			// change fileDir from gralde to make it downloadable by gradle
			if isGradleDir {
				removedEncDir := removeEncryptPath(fileDir)
				newFileDir := folderBuilder(removedEncDir)
				logger.Debug("newFileDir :: ", newFileDir)
				fileDir = newFileDir
			} else {
				removeDir := path.Join(getHomeDir(), configLocation)
				splitRemoveDir := strings.Split(removeDir, "/")
				lengthRemoveDir := len(splitRemoveDir)

				finalPath := strings.Split(fileDir, "/")
				// TODO :: refactor this so we don't need to hardcode +2
				fileDir = path.Join(finalPath[lengthRemoveDir+2:]...)
			}

			logger.Debug("fileDir final :: ", fileDir)

			err := addFileToS3(s, buffer, contentLength, fileDir, s3Bucket)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

// Open the given path of file
// and return the size of the file using buffer
// panic if it fails
func getFileSize(fileDir string) ([]byte, int64) {
	var contentLength int64
	// Open the file for use
	file, err := os.Open(fileDir)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Get file size and read the file content into a buffer
	fileInfo, _ := file.Stat()
	contentLength = fileInfo.Size()
	buffer := make([]byte, contentLength)
	file.Read(buffer)

	return buffer, contentLength
}

// AddFileToS3 will upload a single file to S3, it will require a pre-built aws session
// and will set file info like content type and encryption on the uploaded file.
// fileDir would be the path of S3 to place the file
func addFileToS3(s *session.Session, buffer []byte, contentLength int64, fileDir string, s3Bucket string) error {

	// Config settings: this is where you choose the bucket, filename, content-type etc.
	// of the file you're uploading.
	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(s3Bucket),
		Key:                  aws.String(fileDir),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(contentLength),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})

	return err
}
