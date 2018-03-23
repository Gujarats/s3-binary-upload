package main

import (
	"log"
	"testing"
)

func TestDirBuilderForS3(t *testing.T) {
	testObjects := []struct {
		isGradleDir bool
		directory   string
		expected    string
	}{
		{
			isGradleDir: false,
			directory:   "/home/gujaratsantana/.s3-binary-upload/artifactory/ext-release-local/com/google/api/services/now/google-api-services-now-v1-rev20150528-1.20.0/1.0/google-api-services-now-v1-rev20150528-1.20.0-1.0.jar",
			expected:    "com/google/api/services/now/google-api-services-now-v1-rev20150528-1.20.0/1.0/google-api-services-now-v1-rev20150528-1.20.0-1.0.jar",
		},

		{
			isGradleDir: true,
			directory:   "/home/gujaratsantana/.gradle/caches/modules-2/files-2.1/com.traveloka.common/accessor/1.0.1/5584950fb406c4c44d8c8b5bb0af9d77db4a72a9/accessor-1.0.1.jar",
			expected:    "com/traveloka/common/accessor/1.0.1/accessor-1.0.1.jar",
		},
	}

	for _, testObject := range testObjects {
		result := dirBuilderForS3(testObject.isGradleDir, testObject.directory)
		if result != testObject.expected {
			log.Fatalf("Fail expected = %+v, result = %+v\n", testObject.expected, result)
		}
	}
}
