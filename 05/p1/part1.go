package p1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MrShanks/advent2025/utils"
)

type Interval struct {
	Min int
	Max int
}

func Solve(filepath string) int {
	f, scanner := utils.ReadInput(filepath)
	defer f.Close()

	intervals := make([]Interval, 0)

	for scanner.Scan() {
		raw := scanner.Text()
		if raw == "" {
			scanner.Text()
			break
		}

		minmax := strings.Split(raw, "-")

		min, err := strconv.Atoi(minmax[0])
		if err != nil {
			fmt.Printf("Cannot convert to int: %v", err)
		}

		max, err := strconv.Atoi(minmax[1])
		if err != nil {
			fmt.Printf("Cannot convert to int: %v", err)
		}

		r := Interval{Min: min, Max: max}
		intervals = append(intervals, r)
	}

	ingredients := make([]int, 0)

	for scanner.Scan() {
		raw := scanner.Text()

		ing, err := strconv.Atoi(raw)
		if err != nil {
			fmt.Printf("Cannot convert to int: %v", err)
		}

		ingredients = append(ingredients, ing)
	}

	return calculate(intervals, ingredients)
}

func calculate(intervals []Interval, ingredients []int) int {
	fresh := 0
	for _, ing := range ingredients {
		for _, interv := range intervals {
			if ing >= interv.Min && ing <= interv.Max {
				fresh++
				break
			}
		}
	}

	return fresh
}
