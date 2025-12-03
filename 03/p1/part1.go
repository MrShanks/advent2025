package p1

import (
	"fmt"
	"strconv"

	"github.com/MrShanks/advent2025/utils"
)

func calculateJoultage(bank []int) int {
	max1, max2 := bank[0], bank[len(bank)-1]
	l := 0
	for i := range len(bank) - 1 {
		if bank[i] > max1 {
			max1 = bank[i]
			l = i
		}
	}

	for i := len(bank) - 1; i > l; i-- {
		if bank[i] > max2 {
			max2 = bank[i]
		}
	}

	jstring := fmt.Sprintf("%d%d", max1, max2)
	joultage, err := strconv.Atoi(jstring)
	if err != nil {
		fmt.Println(err)
	}

	return joultage
}

func Solve(filepath string) int {
	f, scanner := utils.ReadInput(filepath)
	defer f.Close()

	counter := 0
	for scanner.Scan() {
		raw := scanner.Text()

		bank := []int{}
		for _, r := range raw {

			battery, err := strconv.Atoi(string(r))
			if err != nil {
				fmt.Println(err)
			}

			bank = append(bank, battery)
		}

		counter += calculateJoultage(bank)
	}

	return counter
}
