package main

import "strings"

func removeEncPath(pathFile string) string {

	var result string
	splitPathFiles := strings.Split(pathFile, backslash)
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
