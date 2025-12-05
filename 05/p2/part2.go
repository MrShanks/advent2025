package p2

import (
	"fmt"
	"sort"
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

	return calculate(intervals)
}

func calculate(intervals []Interval) int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Min < intervals[j].Min
	})

	condensed := []Interval{intervals[0]}

	for _, curr := range intervals[1:] {
		last := &condensed[len(condensed)-1]
		if curr.Min <= last.Max {
			if curr.Max > last.Max {
				last.Max = curr.Max
			}
		} else {
			condensed = append(condensed, curr)
		}
	}

	count := 0
	for _, interval := range condensed {
		count += interval.Max + 1 - interval.Min
	}

	return count
}
