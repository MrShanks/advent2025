package p1

import (
	"github.com/MrShanks/advent2025/utils"
)

func Solve(filepath string) int {
	f, scanner := utils.ReadInput(filepath)
	defer f.Close()

	matrix := make([][]string, 0)
	for scanner.Scan() {
		rawText := scanner.Text()
		row := make([]string, 0)
		for _, l := range rawText {
			row = append(row, string(l))
		}
		matrix = append(matrix, row)
	}

	// Find the starting position 'S'
	activeBeams := make(map[int]bool)

	for _, row := range matrix {
		for c, char := range row {
			if char == "S" {
				activeBeams[c] = true
				break
			}
		}
		if len(activeBeams) > 0 {
			break
		}
	}

	height := len(matrix)

	counter := 0
	for r := range height {
		nextBeams := make(map[int]bool)

		for col := range activeBeams {

			char := matrix[r][col]

			if char == "^" {
				// Splitter hit, increment count and split
				counter++
				nextBeams[col-1] = true
				nextBeams[col+1] = true
			} else {
				// Empty space, beam continues straight down
				nextBeams[col] = true
			}
		}

		activeBeams = nextBeams
	}

	return counter
}
