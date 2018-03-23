package main

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	gradleCacheDir = "/.gradle/caches/modules-2/files-2.1"
)

type Config struct {
	S3Buckets []string `viper:"s3Buckets"`
	S3Bucket  string   `viper:"s3Bucket"`
	Region    string   `viper:"region"`
	Profile   string   `viper:"profile"`

	GradleCacheDir      []string `viper:"gradle"`
	ArtfactsDirectories []string `viper:"artfactsDirectories"`

	UploadArtifacs    bool     `viper:"uploadArtifacs"`
	DownloadArtifacs  bool     `viper:"downloadArtifacs"`
	ArtifactsLocation string   `viper:"artifactsLocation"`
	LinkArtifacts     []string `viper:"linkArtifacts"`

	// authentication for the provided link artifacts
	Username string `viper:"username"`
	Password string `viper:"password"`
}

func getConfig() *Config {
	viper.SetConfigName("config")                  // name of config file (without extension)
	viper.AddConfigPath("$HOME/.s3-binary-upload") // call multiple times to add many search paths
	viper.AddConfigPath(".")                       // optionally look for config in the working directory
	err := viper.ReadInConfig()                    // Find and read the config file
	if err != nil {                                // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file must exist in ~/.s3-binary-upload/config.yaml: %s \n", err))
	}

	// read the config file to struct
	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		panic(fmt.Errorf("Fatal error unmarhsal config struct : %s \n", err))
	}

	return config
}
