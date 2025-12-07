package p2

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/MrShanks/advent2025/utils"
)

func Solve(filepath string) int {
	f, scanner := utils.ReadInput(filepath)
	defer f.Close()

	var rawLines []string
	cols := 0

	for scanner.Scan() {
		line := scanner.Text()
		rawLines = append(rawLines, line)
		cols = len(line)
	}

	matrix := make([][]rune, len(rawLines))
	for i, line := range rawLines {
		matrix[i] = []rune(line)
		fmt.Println(string(matrix[i]))
	}

	var currentBlockNumbers []int
	var currentOp string

	tot := 0
	rows := len(matrix)
	opRowIdx := rows - 1

	processBlock := func() {
		blockTotal := 0
		if currentOp == "*" {
			blockTotal = 1
		}

		for _, num := range currentBlockNumbers {
			switch currentOp {
			case "+":
				blockTotal += num
			case "*":
				blockTotal *= num
			}
		}

		tot += blockTotal

		// reset
		currentBlockNumbers = nil
		currentOp = ""
	}

	for col := range cols {
		isSeparator := true
		for r := range rows {
			if matrix[r][col] != ' ' {
				isSeparator = false
				break
			}
		}

		// when we hit a separator we can start processing
		if isSeparator {
			processBlock()
			continue
		}

		digitStr := ""

		for r := range opRowIdx {
			char := matrix[r][col]
			if unicode.IsDigit(char) {
				digitStr += string(char)
			}
		}
		fmt.Println(string(digitStr))

		if len(digitStr) > 0 {
			num, err := strconv.Atoi(digitStr)
			if err == nil {
				currentBlockNumbers = append(currentBlockNumbers, num)
			}
		}

		opChar := matrix[opRowIdx][col]
		if opChar == '+' || opChar == '*' {
			currentOp = string(opChar)
		}
	}

	processBlock()

	return tot
}
