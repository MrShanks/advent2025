package main

import (
	"fmt"

	"github.com/MrShanks/advent2025/08/p1"
	"github.com/MrShanks/advent2025/08/p2"
)

func main() {
	fmt.Printf("Solution part1: %d\n", p1.Solve("input.txt", 1000))
	fmt.Printf("Solution part2: %d\n", p2.Solve("input.txt"))
}
