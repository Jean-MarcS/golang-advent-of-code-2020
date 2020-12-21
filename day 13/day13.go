package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getDeparture(timeStamp int, busID int) int {
	return ((timeStamp / busID) + 1) * busID
}

func part1(timeStamp int, busLines []int) int {

	min := timeStamp
	goodBus := 0
	for _, busID := range busLines {
		if busID != 0 {
			departure := getDeparture(timeStamp, busID)

			if (departure - timeStamp) < min {
				min = departure - timeStamp
				goodBus = busID
			}

		}
	}
	return min * goodBus
}

func part2(busLines []int) int {
	// All numbers are prime, that's not a coincidence...
	cTS := 0
	step := 1

	for posi, currentBus := range busLines {
		if currentBus != 0 {
			// Find the first occurence that matches with current bus.
			for (cTS+posi)%currentBus != 0 {
				cTS += step
			}
			// Now we know the next step, so let's change it
			step = step * currentBus
		}
	}

	return cTS
}

func main() {
	file, err := os.Open("day13.data")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	var timeStamp int
	var busLinesStr []string
	busLines := []int{}
	scanner.Scan()
	line := scanner.Text()
	timeStamp, _ = strconv.Atoi(line)
	scanner.Scan()
	line = scanner.Text()
	line = strings.ReplaceAll(line, "x", "0")
	busLinesStr = strings.Split(line, ",")

	for _, val := range busLinesStr {
		current, _ := strconv.Atoi(val)
		busLines = append(busLines, current)
	}

	fmt.Println("Part 1 :", part1(timeStamp, busLines))
	fmt.Println("Part 2 :", part2(busLines))
}
