package p1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MrShanks/advent2025/utils"
)

func invalid(current string) bool {
	return current[:len(current)/2] == current[len(current)/2:]
}

func Solve(filepath string) int {
	f, scanner := utils.ReadInput(filepath)
	defer f.Close()

	counter := 0
	for scanner.Scan() {
		raw := scanner.Text()
		ranges := strings.SplitSeq(raw, ",")

		for r := range ranges {

			parts := strings.Split(r, "-")
			start, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Println(err)
			}

			end, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println(err)
			}

			for current := start; current <= end; current++ {
				strC := strconv.Itoa(current)

				if invalid(strC) {
					counter += current
				}

			}
		}
	}
	return counter
}
