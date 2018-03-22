package main

import (
	"bytes"
	"log"
	"net/http"
	"os"

	"github.com/Gujarats/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// pass buckets to uploader to start uploading
func uploadFromGradleCache(s *session.Session, buckets []string, artifacts map[string][]string) {
	for _, artifact := range artifacts {
		logger.Debug("artifact :: ", artifact)
		s3Bucket := nextS3(buckets)
		for _, fileDir := range artifact {
			err := AddFileToS3forGradle(s, fileDir, s3Bucket)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

// pass buckets to uploader to start uploading
func upload(s *session.Session, buckets []string, artifacts map[string][]string) {
	for _, artifact := range artifacts {
		logger.Debug("artifact :: ", artifact)
		s3Bucket := nextS3(buckets)
		for _, fileDir := range artifact {
			err := AddFileToS3(s, fileDir, s3Bucket)
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

	// Config settings: this is where you choose the bucket, filename, content-type etc.
	// of the file you're uploading.

	_, err = s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(s3Bucket),
		Key:                  aws.String(fileDir),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})

	return err
}

// AddFileToS3 will upload a single file to S3, it will require a pre-built aws session
// and will set file info like content type and encryption on the uploaded file.
func AddFileToS3forGradle(s *session.Session, fileDir string, s3Bucket string) error {

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
