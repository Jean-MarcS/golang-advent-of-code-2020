package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(line string) int {
	row, _ := strconv.ParseInt(strings.Replace(strings.Replace(line[0:7], "F", "0", -1), "B", "1", -1), 2, 64)
	col, _ := strconv.ParseInt(strings.Replace(strings.Replace(line[7:], "R", "1", -1), "L", "0", -1), 2, 64)
	return int((row * 8) + col)
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

	for scanner.Scan() {
		line := scanner.Text()
		currentSeatID := part1(line)
		if currentSeatID > maxSeatID {
			maxSeatID = currentSeatID
		}
	}

	fmt.Println("Part 1 : ", maxSeatID)

}
