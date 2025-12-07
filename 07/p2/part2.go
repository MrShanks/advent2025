package p2

import (
	"github.com/MrShanks/advent2025/utils"
)

func calculate(matrix [][]string) int {
	height := len(matrix)
	if height == 0 {
		return 0
	}
	width := len(matrix[0])

	currentCounts := make([]int, width)
	nextCounts := make([]int, width)

	// Find Start 'S' and initialize
	startRow := 0
	for r, row := range matrix {
		found := false
		for c, char := range row {
			if char == "S" {
				currentCounts[c] = 1
				startRow = r
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	total := 0

	// Process row by row
	for r := startRow; r < height; r++ {
		// Reset next row buffer to 0
		for i := range nextCounts {
			nextCounts[i] = 0
		}

		activeCount := 0
		for c := range width {
			count := currentCounts[c]
			if count == 0 {
				continue
			}
			activeCount++

			char := matrix[r][c]

			if char == "^" {
				// Split Left
				nextCounts[c-1] += count

				// Split Right
				nextCounts[c+1] += count
			} else {
				// Fall straight down
				nextCounts[c] += count
			}
		}

		// Swap buffers
		currentCounts, nextCounts = nextCounts, currentCounts
	}

	// Sum up any timelines that reached the bottom successfully
	for _, count := range currentCounts {
		total += count
	}

	return total
}

func Solve(filepath string) int {
	f, scanner := utils.ReadInput(filepath)
	defer f.Close()

	// Parse grid first to get dimensions
	matrix := make([][]string, 0)
	for scanner.Scan() {
		rawText := scanner.Text()
		row := make([]string, 0)
		for _, l := range rawText {
			row = append(row, string(l))
		}
		matrix = append(matrix, row)
	}

	return calculate(matrix)

}
