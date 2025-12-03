package p2

import (
	"fmt"
	"strconv"

	"github.com/MrShanks/advent2025/utils"
)

func calculateJoultage(bank string, k int) string {
	n := len(bank)
	if n <= k {
		return bank
	}

	dropCount := n - k

	stack := make([]byte, 0, n)

	for i := range n {
		digit := bank[i]

		for len(stack) > 0 && dropCount > 0 && stack[len(stack)-1] < digit {
			stack = stack[:len(stack)-1]
			dropCount--
		}

		stack = append(stack, digit)
	}
	return string(stack[:k])
}

func Solve(filepath string) int64 {
	f, scanner := utils.ReadInput(filepath)
	defer f.Close()

	var counter int64
	for scanner.Scan() {
		raw := scanner.Text()

		joultage := calculateJoultage(raw, 12)

		val, err := strconv.ParseInt(joultage, 10, 64)
		if err != nil {
			fmt.Printf("Error parsing number %s: %v\n", joultage, err)
			continue
		}

		counter += val
	}

	return counter
}
