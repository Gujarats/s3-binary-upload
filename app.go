package main

import (
	"fmt"
	"log"
	"path"

	"github.com/Gujarats/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	config := getConfig()
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(config.Region),
			Credentials: credentials.NewCredentials(
				&credentials.SharedCredentialsProvider{
					Profile: config.Profile,
				},
			),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	if config.DownloadArtifacs {
		Downloader(config.LinkArtifacts, config.Username, config.Password)
	}

	if config.UploadArtifacs {
		artifactsDirectories, isGradleDir := getArtifactsDir(config)

		// get package name
		fmt.Print("\nEnter your package = ")
		var packageName string
		_, err = fmt.Scan(&packageName)
		if err != nil {
			log.Fatal(err)
		}

		var packageNames []string
		var currentDir string
		for _, artifactsDir := range artifactsDirectories {
			// list all the directory names
			currentDir = path.Join(getHomeDir(), config.ArtifactsLocation, artifactsDir)
			logger.Debug("currentDir :: ", currentDir)

			result := runCommand("ls", currentDir, false)

			// create directory builder here
			fullDir := directoryBuilder(currentDir, result)
			packageNames = append(packageNames, fullDir...)
		}

		logger.Debug("packageNames :: ", packageNames)

		// store arifact jar & pom
		// key = artifact name prefix without extenstion like com.traveloka.common/accessor-1.0.2
		// []string all files dir
		artifacts := make(map[string][]string)

		// get specific directory for scanning artifact
		//packages := filterDir(firtsDirPackageNames, packageName)
		//logger.Debug("packages result :: ", packages)

		for _, pack := range packageNames {
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
		upload(sess, config, buckets, artifacts, isGradleDir)
	}
}

func getArtifactsDir(config *Config) ([]string, bool) {
	fromGradle := false
	if len(config.ArtfactsDirectories) > 0 {
		return config.ArtfactsDirectories, fromGradle
	} else if len(config.GradleCacheDir) > 0 {
		fromGradle = true
		return config.GradleCacheDir, fromGradle
	}

	return []string{}, fromGradle
}
