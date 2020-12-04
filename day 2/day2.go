package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func testPart1(min int, max int, c string, pass string) int {
	nb := strings.Count(pass, c)

	if nb >= min && nb <= max {
		return 1
	} else {
		return 0
	}
}

func testPart2(i1 int, i2 int, c string, pass string) int {
	total := 0

	if pass[i1-1:i1] == c {
		total = 1
	}

	if pass[i2-1:i2] == c {
		total++
	}

	if total == 1 {
		return 1
	} else {
		return 0
	}
}

func checkPass(line string, part int) int {

	i := strings.Index(line, "-")
	n1, _ := strconv.Atoi(line[:i])
	line = line[i+1:]

	i = strings.Index(line, " ")
	n2, _ := strconv.Atoi(line[:i])
	c := line[i+1 : i+2]
	pass := line[i+4:]

	if part == 1 {
		return testPart1(n1, n2, c, pass)
	} else {
		return testPart2(n1, n2, c, pass)
	}

}

func main() {
	file, err := os.Open("day2.data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalP1 := 0
	totalP2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		totalP1 += checkPass(line, 1)
		totalP2 += checkPass(line, 2)
	}

	fmt.Println("Total part1 :", totalP1)
	fmt.Println("Total part2 :", totalP2)

}
