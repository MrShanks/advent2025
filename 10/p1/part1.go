package p1

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/MrShanks/advent2025/utils"
)

func calculate(line string) int {
	lightRegex := regexp.MustCompile(`\[([.#]+)\]`)
	matches := lightRegex.FindStringSubmatch(line)
	if matches == nil {
		return 0
	}

	numLights := len(matches[1])
	var target uint64
	for i, ch := range matches[1] {
		if ch == '#' {
			target |= (1 << i)
		}
	}

	btnRegex := regexp.MustCompile(`\(([\d,]+)\)`)
	btnMatches := btnRegex.FindAllStringSubmatch(line, -1)

	buttons := make([]uint64, len(btnMatches))
	for i, m := range btnMatches {
		parts := strings.SplitSeq(m[1], ",")
		for p := range parts {
			bitIdx, _ := strconv.Atoi(p)
			buttons[i] |= (1 << bitIdx)
		}
	}

	// Brute Force all subsets of buttons
	// Since order doesn't matter (A then B is same as B then A),
	// and pressing twice cancels out, we only need to check each subset once.
	minPresses := math.MaxInt
	limit := 1 << len(buttons) // 2^N combinations

	for mask := range limit {
		var currentLights uint64
		presses := 0

		for i, btnPress := range buttons {
			if (mask>>i)&1 == 1 {
				currentLights ^= btnPress
				presses++
			}
		}

		visualizeStep(mask, currentLights, target, numLights, buttons)

		if currentLights == target {
			if presses < minPresses {
				minPresses = presses
			}
		}
	}

	if minPresses == math.MaxInt {
		return 0
	}
	return minPresses
}

func Solve(filepath string) int {
	f, scanner := utils.ReadInput(filepath)
	defer f.Close()

	total := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			total += calculate(line)
		}
	}

	return total
}

func visualizeStep(mask int, currentLights, target uint64, numLights int, buttons []uint64) {
	fmt.Print("\033[H\033[2J")

	fmt.Printf("Checking Strategy (Mask %d) \n", mask)

	fmt.Print("Buttons Pressed: ")
	pressed := []string{}
	for i := range buttons {
		if (mask>>i)&1 == 1 {
			pressed = append(pressed, fmt.Sprintf("Btn%d", i))
		}
	}
	if len(pressed) == 0 {
		fmt.Print("(None)")
	} else {
		fmt.Print(strings.Join(pressed, ", "))
	}
	fmt.Printf("\n\n")

	printBits := func(label string, val uint64) {
		fmt.Printf("%s [", label)
		for i := range numLights {
			if (val>>i)&1 == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("]")
	}

	printBits("Current State: ", currentLights)
	printBits("Target State:  ", target)

	if currentLights == target {
		fmt.Println("\nSTATUS: *** MATCH FOUND! ***")
	} else {
		fmt.Println("\nSTATUS: Mismatch...")
	}

	time.Sleep(1500 * time.Millisecond)
}
