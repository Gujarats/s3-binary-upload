package main

import (
	"regexp"
	"strings"
)

const (
	backslash = "/"
	dot       = "."
)

// get root folder and subfolders from given path of artifact
func getFolder(source string) []string {
	var result []string
	re := regexp.MustCompile("\\d")
	splitBackslashes := strings.Split(source, backslash)
	for _, splitBackslash := range splitBackslashes {
		if re.MatchString(splitBackslash) {
			result = append(result, splitBackslash)
			continue
		}
		splitDots := strings.Split(splitBackslash, dot)
		for _, splitDot := range splitDots {
			result = append(result, splitDot)
		}

	}
	return result
}
