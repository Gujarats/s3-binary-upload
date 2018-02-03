package main

import "testing"

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
