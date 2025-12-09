package p1

import (
	"strconv"
	"strings"

	"github.com/MrShanks/advent2025/utils"
)

type Point struct {
	x, y int
}

func Solve(filepath string) int {
	f, scanner := utils.ReadInput(filepath)
	defer f.Close()

	points := []Point{}
	for scanner.Scan() {
		raw := scanner.Text()
		parts := strings.Split(raw, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		points = append(points, Point{x: x, y: y})
	}

	var maxArea int
	for i, p1 := range points {
		if i+1 >= len(points)-1 {
			break
		}

		for _, p2 := range points[i+1:] {
			x, y := 0, 0
			x = abs(p1.x - p2.x)
			y = abs(p1.y - p2.y)

			current := (x + 1) * (y + 1)
			maxArea = max(maxArea, current)
		}
	}

	return maxArea
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
