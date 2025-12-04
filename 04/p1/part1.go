package p1

import (
	"github.com/MrShanks/advent2025/utils"
)

func Solve(filepath string) int {
	f, scanner := utils.ReadInput(filepath)
	defer f.Close()

	var matrix [][]string
	for scanner.Scan() {
		raw := scanner.Text()
		var row []string
		for _, r := range raw {
			row = append(row, string(r))
		}
		matrix = append(matrix, row)
	}

	rows := len(matrix)
	cols := len(matrix[0])
	counter := 0

	for r := range rows {
		for c := range cols {
			neighbour := 0
			for _, dr := range []int{-1, 0, 1} {
				for _, dc := range []int{-1, 0, 1} {
					posr, posc := r+dr, c+dc
					if posr >= 0 && posr < rows && posc >= 0 && posc < cols && matrix[posr][posc] == "@" {
						neighbour += 1
					}
				}
			}
			if matrix[r][c] == "@" && neighbour < 5 {
				counter += 1
			}
		}
	}

	return counter
}
