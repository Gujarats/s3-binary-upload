package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/Gujarats/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var wg sync.WaitGroup

// pass buckets to uploader to start uploading
func upload(s *session.Session, config *Config, buckets []string, artifacts map[string][]string, isGradleDir bool) {
	for _, artifact := range artifacts {
		s3Bucket := nextS3(buckets)
		for _, fileDir := range artifact {
			buffer, contentLength := getFileSize(fileDir)
			dirForS3 := dirBuilderForS3(isGradleDir, fileDir)

			wg.Add(1)
			go func(s *session.Session, buffer []byte, contentLength int64, fileDir string, s3Bucket string) {

				err := addFileToS3(s, buffer, contentLength, dirForS3, s3Bucket)
				if err != nil {
					log.Fatal(err)
				}
				wg.Done()
			}(s, buffer, contentLength, dirForS3, s3Bucket)
		}
	}

	wg.Wait()
}

// This Will create directory for artifacts so it can be downloaded by gradle
// by following maven repository convention
// directory is the current dir for artifacts in local disk
func dirBuilderForS3(isGradleDir bool, directory string) string {
	var result string

	// change fileDir from gralde to make it downloadable by gradle
	if isGradleDir {
		removeDir := path.Join(getHomeDir(), gradleCacheDir)
		splitRemoveDir := strings.Split(removeDir, "/")
		lengthRemoveDir := len(splitRemoveDir)

		fullPath := strings.Split(directory, "/")
		dirWithoutGradleCache := path.Join(fullPath[lengthRemoveDir:]...)

		removedEncDir := removeEncryptPath(dirWithoutGradleCache)
		newFileDir := folderBuilder(removedEncDir)
		result = newFileDir
	} else {
		removeDir := path.Join(getHomeDir(), configLocation)
		splitRemoveDir := strings.Split(removeDir, "/")
		lengthRemoveDir := len(splitRemoveDir)

		finalPath := strings.Split(directory, "/")
		// TODO :: refactor this so we don't need to hardcode +2
		result = path.Join(finalPath[lengthRemoveDir+2:]...)
	}

	return result
}

// Open the given path of file
// and return the size of the file using buffer
// panic if it fails
func getFileSize(fileDir string) ([]byte, int64) {
	var contentLength int64
	// Open the file for use
	file, err := os.Open(fileDir)
	if err != nil {
		logger.Debug("error :: ", err)
		log.Fatal(err)
	}
	//defer file.Close()

	// Get file size and read the file content into a buffer
	fileInfo, _ := file.Stat()
	contentLength = fileInfo.Size()
	buffer := make([]byte, contentLength)
	file.Read(buffer)
	file.Close()

	return buffer, contentLength
}

// AddFileToS3 will upload a single file to S3, it will require a pre-built aws session
// and will set file info like content type and encryption on the uploaded file.
// fileDir would be the path of S3 to place the file
func addFileToS3(s *session.Session, buffer []byte, contentLength int64, fileDir string, s3Bucket string) error {

	object, err := s3.New(s).GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(fileDir),
	})

	//upload new file if objects is not exist
	if err != nil || *object.ContentLength <= 0 {

		logger.Debug("fileDir upload :: ", fileDir)

		// upload the file with given buffer and key
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
	} else {
		logger.Debug("fileDir skipped :: ", fileDir)
	}

	return nil
}
