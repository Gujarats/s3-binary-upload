package main

import (
	"reflect"
	"testing"
)

func TestFolderBuilder(t *testing.T) {
	testObjects := []struct {
		source   string
		expected string
	}{
		{
			source:   "com.helloworld.common.gradle/3.01/java.jar",
			expected: "com/helloworld/common/gradle/3.01/java.jar",
		},
		{
			source:   "com.helloworld.common.gradle/3.01/java.jar",
			expected: "com/helloworld/common/gradle/3.01/java.jar",
		},
		{
			source:   "com.helloworld.common.gradle/3.01.Final/java.jar",
			expected: "com/helloworld/common/gradle/3.01.Final/java.jar",
		},
	}

	for _, testObject := range testObjects {
		actual := folderBuilder(testObject.source)
		if actual != testObject.expected {
			t.Errorf("expected = %+v, actual = %+v\n", testObject.expected, actual)
		}
	}
}

func TestRemoveDotToSlash(t *testing.T) {
	testObjects := []struct {
		source   string
		expected []string
	}{
		//{
		//	source:   "com.helloworld.common.gradle",
		//	expected: []string{"com", "helloworld", "common", "gradle"},
		//},

		//{
		//	source:   "com.helloworld.common.gradle/library",
		//	expected: []string{"com", "helloworld", "common", "gradle", "library"},
		//},

		//{
		//	source:   "com.hello-world.common.gradle/library/java-plugin/4.10.0",
		//	expected: []string{"com", "hello-world", "common", "gradle", "library", "java-plugin", "4.10.0"},
		//},

		//{
		//	source:   "com.hello-world.p2p.gradle/library/java-plugin/4.10.0",
		//	expected: []string{"com", "hello-world", "p2p", "gradle", "library", "java-plugin", "4.10.0"},
		//},
		{
			source:   "com.hello-world.p2p.gradle/library/4.1.13.Final/checkout",
			expected: []string{"com", "hello-world", "p2p", "gradle", "library", "4.1.13.Final", "checkout"},
		},
	}

	for _, testObject := range testObjects {
		actual := removeDotToSlash(testObject.source)
		if !reflect.DeepEqual(actual, testObject.expected) {
			t.Errorf("expected = %+v, actual = %+v\n", testObject.expected, actual)
		}
	}
}

func TestRemoveEncPath(t *testing.T) {
	testObjects := []struct {
		source   string
		expected string
	}{
		{
			source:   "com.helloworld.common/accessor/1.0.0/16aafadb75b1ad0eacf6e3fef68320e6502df136/file.pom",
			expected: "com.helloworld.common/accessor/1.0.0/file.pom",
		},

		{
			source:   "com.helloworld.common/accessor/1.0.0/16aafadb75b1ad0eacf6e3fef68320e6502df136/file.pom",
			expected: "com.helloworld.common/accessor/1.0.0/file.pom",
		},
	}

	for _, testObject := range testObjects {
		actual := removeEncryptPath(testObject.source)
		if actual != testObject.expected {
			t.Errorf("expected = %+v, actual = %+v\n", testObject.expected, actual)
		}
	}
}
