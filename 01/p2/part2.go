package p2

import (
	"strconv"

	"github.com/MrShanks/advent2025/utils"
)

func Solve(filepath string) int {
	f, scanner := utils.ReadInput(filepath)
	defer f.Close()

	current := 50
	counter := 0

	for scanner.Scan() {
		raw := scanner.Text()
		abs, _ := strconv.Atoi(raw[1:])

		rotations := 0
		number := 0

		sign := raw[0]
		if sign == 'L' {
			number = -abs

			rotations = (current + number - 100) / -100

			if current == 0 {
				rotations--
			}
		} else {
			number = abs
			rotations = (current + number) / 100
		}

		counter += rotations

		current = (current + number) % 100
		if current < 0 {
			current += 100
		}
	}

	return counter
}
