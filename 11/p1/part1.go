package p1

import (
	"fmt"

	"github.com/MrShanks/advent2025/utils"
)

func Solve(filepath string) int {
	f, scanner := utils.ReadInput(filepath)
	defer f.Close()

	counter := 0
	for scanner.Scan() {
		raw := scanner.Text()

		fmt.Println(raw)
	}

	return counter
}
