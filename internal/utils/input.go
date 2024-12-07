package utils

import (
	"os"
	"strings"
)

func ReadInputFile(fileName string) []byte {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return data
}

func ReadInputFileByLine(fileName string, cb func(i int, row string)) {
	data := ReadInputFile(fileName)
	for i, row := range strings.Split(string(data), "\n") {
		cb(i, row)
	}
}
