package main

import "testing"

func TestGetFolder(t *testing.T) {
	testObjects := []struct {
		source   string
		expected []string
		err      error
	}{
		{
			source:   "com.helloworld.common.gradle",
			expected: []string{"com", "helloworld", "common", "gradle"},
			err:      nil,
		},

		{
			source:   "com.helloworld.common.gradle/library",
			expected: []string{"com", "helloworld", "common", "gradle", "library"},
			err:      nil,
		},

		{
			source:   "com.hello-world.common.gradle/library/java-plugin/4.10.0",
			expected: []string{"com", "hello-world", "common", "gradle", "library", "java-plugin", "4.10.0"},
			err:      nil,
		},
	}

	for _, testObject := range testObjects {
		actual := getFolder(testObject.source)
		if !sliceString(testObject.expected, actual) {
			t.Errorf("expected = %+v, actual = %+v\n", testObject.expected, actual)
		}
	}
}
func sliceString(a, b []string) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
