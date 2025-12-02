package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInput(filepath string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)

	return file, scanner
}
