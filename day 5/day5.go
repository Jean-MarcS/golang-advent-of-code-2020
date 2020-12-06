package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1(line string) int {
	row, _ := strconv.ParseInt(strings.Replace(strings.Replace(line[0:7], "F", "0", -1), "B", "1", -1), 2, 64)
	col, _ := strconv.ParseInt(strings.Replace(strings.Replace(line[7:], "R", "1", -1), "L", "0", -1), 2, 64)
	return int((row * 8) + col)
}

func part2(seats []int) int {
	sort.Ints(seats)
	for i := 1; i < len(seats); i++ {
		if seats[i-1]+1 != seats[i] {
			return seats[i-1] + 1
		}
	}
	return -1
}

func main() {

	file, err := os.Open("day5.data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var maxSeatID int
	maxSeatID = 0
	var seats []int

	for scanner.Scan() {
		line := scanner.Text()
		currentSeatID := part1(line)
		if currentSeatID > maxSeatID {
			maxSeatID = currentSeatID
		}
		seats = append(seats, currentSeatID)
	}

	fmt.Println("Part 1 : ", maxSeatID)
	fmt.Println("Part 2 : ", part2(seats))

}
