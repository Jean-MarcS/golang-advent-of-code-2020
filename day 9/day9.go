package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const totalSize = 1000
const preambleSize = 25

func part1(currentPreamble []int, current int) int {
	found := false
	cpt1 := 0
	for cpt1 < preambleSize-1 {
		cpt2 := cpt1 + 1
		for cpt2 < preambleSize {
			if current == currentPreamble[cpt1]+currentPreamble[cpt2] {
				found = true
			}
			cpt2++
		}
		cpt1++
	}
	if found {
		return 0
	}
	return current

}

func part2(lines [totalSize]int, weakness int) int {
	cpt1 := 0
	cpt2 := 0
	found := false
	for cpt1 < totalSize-1 && !found {
		cpt2 = cpt1
		currentSum := 0
		for cpt2 < totalSize && !found {
			currentSum += lines[cpt2]
			if currentSum > weakness {
				break
			}
			if currentSum == weakness {
				found = true
				break
			}
			cpt2++
		}
		if found {
			break
		}
		cpt1++
	}

	min := weakness
	max := 0
	for i := cpt1; i <= cpt2; i++ {
		if lines[i] < min {
			min = lines[i]
		}
		if lines[i] > max {
			max = lines[i]
		}

	}
	return min + max
}

func main() {
	file, err := os.Open("day9.data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cpt := 0
	var lines [totalSize]int

	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		lines[cpt] = val
		cpt++
	}

	cpt = preambleSize
	weakness := 0
	currentPreamble := lines[0:preambleSize]

	for cpt < totalSize && weakness == 0 {
		weakness = part1(currentPreamble, lines[cpt])
		currentPreamble = append(currentPreamble[1:], []int{lines[cpt]}...)
		cpt++
	}

	fmt.Println("Part 1 :", weakness)
	fmt.Println("Part 2 :", part2(lines, weakness))
}
