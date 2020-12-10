package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instrType struct {
	code string
	val  int
	used bool
}

func part2(instructions []instrType, totalSize int) int {
	pt := 0
	accumulator := 0
	for !instructions[pt].used && pt < totalSize {
		instructions[pt].used = true
		switch instructions[pt].code {
		case "nop":
			pt++
		case "jmp":
			pt += instructions[pt].val
		case "acc":
			accumulator += instructions[pt].val
			pt++

		}
	}
	if pt == totalSize {
		//fmt.Println(pt, accumulator)
		return accumulator
	} else {
		return 0
	}
}

func main() {

	totalSize := 638 // my input size
	instructions := make([]instrType, totalSize+1)
	file, err := os.Open("day8.data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cpt := 0

	for scanner.Scan() {
		line := scanner.Text()
		datas := strings.Split(line, " ")
		val, _ := strconv.Atoi(datas[1])
		instructions[cpt] = instrType{code: datas[0], val: val}
		cpt++
	}

	// Part 1
	pt := 0
	accumulator := 0
	for !instructions[pt].used {
		instructions[pt].used = true
		switch instructions[pt].code {
		case "nop":
			pt++
		case "jmp":
			pt += instructions[pt].val
		case "acc":
			accumulator += instructions[pt].val
			pt++

		}
	}

	fmt.Println("Part 1 : ", accumulator)

	accumulator = 0
	// I know all of the instruction that can be changed (in "used") so let's only bruteforce those
	for pt := range instructions {
		if instructions[pt].used {
			tmpInstructions := make([]instrType, totalSize+1)
			copy(tmpInstructions, instructions) // Make a copy
			for tmpPt := range tmpInstructions {
				// reinit visited
				tmpInstructions[tmpPt].used = false
			}
			if tmpInstructions[pt].code == "nop" {
				tmpInstructions[pt].code = "jmp"
			} else {
				if tmpInstructions[pt].code == "jmp" {
					tmpInstructions[pt].code = "nop"
				}
			}
			test := part2(tmpInstructions, totalSize)
			if test != 0 {
				accumulator = test
			}
		}
	}

	fmt.Println("Part 2 : ", accumulator)

}
