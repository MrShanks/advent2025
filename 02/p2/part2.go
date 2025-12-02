package p2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MrShanks/advent2025/utils"
)

func invalid(current string) bool {
	for i := 2; i <= len(current); i++ {
		if len(current)%i == 0 {
			ok := true
			size := len(current) / i
			// walk j by blocks of size
			for j := 0; j < len(current); j += size {
				// compare to the first block, if they are all alike then we have a winner
				if current[j:j+size] != current[:size] {
					ok = false
					break
				}
			}
			if ok {
				return true
			}
		}
	}
	return false
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
