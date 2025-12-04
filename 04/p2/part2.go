package p2

import (
	"github.com/MrShanks/advent2025/utils"
)

func calculate(matrix [][]string) int {
	rows := len(matrix)
	cols := len(matrix[0])
	counter := 0

	changed := true

	for changed {
		changed = false

		//loop trough each cell
		for r := range rows {
			for c := range cols {

				neighbour := 0
				//loop trough each neighbour
				for _, dr := range []int{-1, 0, 1} {
					for _, dc := range []int{-1, 0, 1} {
						posr, posc := r+dr, c+dc

						//check wew are within boundaries and the cell is a rollpaper
						if posr >= 0 && posr < rows && posc >= 0 && posc < cols && matrix[posr][posc] == "@" {
							neighbour += 1
						}
					}
				}
				//if we have less than 5 neighbours then we can remove the paperoll
				if matrix[r][c] == "@" && neighbour < 5 {
					changed = true
					counter += 1
					matrix[r][c] = "."
				}
			}
		}
	}
	return counter
}

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

	tot := calculate(matrix)

	return tot
}
