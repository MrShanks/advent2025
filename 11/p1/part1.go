package p1

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

	return search(servers, "you")
}

func search(servers map[string][]string, key string) int {
	if key == "out" {
		return 1
	}

	numberOfPaths := 0
	for _, v := range servers[key] {
		numberOfPaths += search(servers, v)
	}
	return numberOfPaths
}
