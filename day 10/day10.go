package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// I failed on part 2 so stole the answer from https://zestedesavoir.com/billets/3723/ladvent-of-code-2020-en-go-jours-6-a-10/#5-jour-10-un-probleme-de-denombrement
// :-)
func part2(input []int) int {
	x := 0                       // Last computed joltage (initialize on the 0j input)
	sum, prev1, prev2 := 1, 0, 0 // Number of arrangements for x, x-1, x-2 jolts
	for _, n := range input {
		for x < n-1 {
			// "Slide the window" until we have the number of arrangements for
			// n-1, n-2, n-3 jolts: fill missing values with 0.
			x++
			prev2, prev1, sum = prev1, sum, 0
		}
		x++
		prev2, prev1, sum = prev1, sum, sum+prev1+prev2
	}
	return sum
}

func main() {
	file, err := os.Open("day10.data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []int
	var diff [4]int

	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, []int{val}...)
	}

	sort.Ints(lines)
	current := 0 // Outlet
	for _, val := range lines {

		thisDiff := val - current
		diff[thisDiff]++
		current = val
	}
	diff[3]++ // My Device

	fmt.Println("Part 1 :", diff[1]*diff[3])
	fmt.Println("Part 2 :", part2(lines))

}
