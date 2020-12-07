package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func initDecl() map[string]int {
	customDecl := make(map[string]int)

	for i := 'a'; i <= 'z'; i++ {
		customDecl[string(i)] = 0
	}

	return customDecl
}

func checkPart1(customDecl map[string]int, line string) map[string]int {
	for _, char := range line {
		customDecl[string(char)] = 1
	}
	return customDecl
}

func checkPart2(customDecl map[string]int, line string) map[string]int {
	for _, char := range line {
		customDecl[string(char)]++
	}
	return customDecl
}

func main() {

	var customDeclP1 map[string]int
	var customDeclP2 map[string]int

	file, err := os.Open("day6.data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	customDeclP1 = initDecl()
	customDeclP2 = initDecl()
	totalYesP1 := 0
	totalYesP2 := 0
	nbPersonsInGroup := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 {
			customDeclP1 = checkPart1(customDeclP1, line)
			customDeclP2 = checkPart2(customDeclP2, line)
			nbPersonsInGroup++
		} else {
			// Add new yes answers
			for _, val := range customDeclP1 {
				totalYesP1 += val
			}
			for _, val := range customDeclP2 {
				if val == nbPersonsInGroup {
					totalYesP2++
				}
			}
			customDeclP1 = initDecl()
			customDeclP2 = initDecl()
			nbPersonsInGroup = 0
		}
	}

	fmt.Println("Part 1 : ", totalYesP1)
	fmt.Println("Part 2 : ", totalYesP2)

}
