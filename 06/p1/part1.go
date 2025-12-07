package p1

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/MrShanks/advent2025/utils"
)

func calculate(cols [][]int, operand []string) int {

	tot := make([]int, 0)

	for i := range len(cols[0]) {
		switch operand[i] {
		case "+":
			{
				count := 0
				for j := range len(cols) {
					count += cols[j][i]
				}
				tot = append(tot, count)
			}
		case "*":
			{
				count := 1
				for j := range len(cols) {
					count *= cols[j][i]
				}
				tot = append(tot, count)
			}
		}
	}

	count := 0
	for _, e := range tot {
		count += e
	}

	return count
}

func Solve(filepath string) int {
	f, scanner := utils.ReadInput(filepath)
	defer f.Close()

	cols := make([][]string, 0)
	for scanner.Scan() {
		raw := scanner.Text()

		numbers := strings.Fields(raw)

		cols = append(cols, numbers)
	}

	//remove operands
	operands := cols[len(cols)-1]
	cols = slices.Delete(cols, len(cols)-1, len(cols))

	columns := [][]int{}
	for _, row := range cols {
		numbers := []int{}
		for _, str := range row {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("failed to conver:", err)
				continue
			}
			numbers = append(numbers, num)
		}
		columns = append(columns, numbers)
	}
	return calculate(columns, operands)
}
