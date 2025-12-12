package p2

import (
	"strings"

	"github.com/MrShanks/advent2025/utils"
)

func Solve(filepath string) int {
	f, scanner := utils.ReadInput(filepath)
	defer f.Close()

	servers := make(map[string][]string)
	for scanner.Scan() {
		raw := scanner.Text()

		parts := strings.Split(raw, ":")
		servers[parts[0]] = strings.Fields(parts[1])
	}

	search := func(start, end string) int {
		memo := make(map[string]int)
		return countPathsMemo(servers, start, end, memo)
	}

	path1 := search("svr", "dac") *
		search("dac", "fft") *
		search("fft", "out")

	path2 := search("svr", "fft") *
		search("fft", "dac") *
		search("dac", "out")

	return path1 + path2
}

// had to add memoization to avoid computing paths multiple times.
func countPathsMemo(servers map[string][]string, current, target string, memo map[string]int) int {
	if current == target {
		return 1
	}

	if val, found := memo[current]; found {
		return val
	}

	totalPaths := 0
	for _, neighbor := range servers[current] {
		totalPaths += countPathsMemo(servers, neighbor, target, memo)
	}

	memo[current] = totalPaths

	return totalPaths
}
