package main

import "testing"

func TestNextS3(t *testing.T) {
	testObjects := []struct {
		source   []string
		expected []string
	}{
		{
			source:   []string{"s1", "s2", "s3"},
			expected: []string{"s1", "s2", "s3", "s1", "s2", "s3", "s1"},
		},
	}

	for _, testObject := range testObjects {
		for _, expected := range testObject.expected {
			actual := nextS3(testObject.source)
			if actual != expected {
				t.Errorf("expected = %+v, actual = %+v\n", expected, actual)
			}
		}
	}
}
