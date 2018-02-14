package main

import (
	"regexp"
	"strings"
)

func folderBuilder(source string) string {
	var result string
	folders := removeDotToSlash(source)
	for index, folder := range folders {
		if index != len(folders)-1 {
			result += folder + "/"
		} else {
			result += folder
		}

	}

	return result
}

func removeDotToSlash(source string) []string {
	var result []string
	re := regexp.MustCompile("^\\d+(\\.\\d+)*([a-zA-Z])*")
	splitSlashes := strings.Split(source, slash)
	for index, splitSlash := range splitSlashes {
		// avoid splitting the dots for the filename
		if index == len(splitSlashes)-1 && len(splitSlashes) != 1 {
			result = append(result, splitSlash)
			continue
		}

		// avoid spliting removing the dot for the number in folder name
		if re.MatchString(splitSlash) {
			result = append(result, splitSlash)
			continue
		}

		splitDots := strings.Split(splitSlash, dot)
		for _, splitDot := range splitDots {
			result = append(result, splitDot)
		}
	}
	return result
}

func removeEncryptPath(pathFile string) string {
	var result string
	splitPathFiles := strings.Split(pathFile, slash)
	for index, splitPathFile := range splitPathFiles {
		if index != (len(splitPathFiles) - 2) {
			if index == 0 {
				result = splitPathFile
			} else {
				result = result + "/" + splitPathFile
			}
		}
	}
	return result
}
