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

func part1(instructions []instrType) int {
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
	return accumulator
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

	fmt.Println("Part 1 : ", part1(instructions))
	fmt.Println("Part 2 : ")

}
