package p1

import (
	"fmt"
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
		number, err := strconv.Atoi(raw[1:])
		if err != nil {
			fmt.Println(err)
		}

		sign := raw[0]
		if sign == 'L' {
			number = number * -1
		}

		//the module makes sure every overflow is correctly reset to the start of dial
		// 103 % 100 = 3 the initial addition of 100 makes sure the are no negative results
		// that would yield a negative module result
		current = (100 + current + number) % 100

		if current == 0 {
			counter++
		}
	}

	return counter
}
