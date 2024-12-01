package utils

import (
	"io"
	"log"
	"os"
)

func FileToString(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input, err := io.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	return string(input)
}
