package main

var index = -1

func nextS3(source []string) string {

	if index >= len(source)-1 {
		index = 0
	} else {
		index++
	}

	return source[index]
}
