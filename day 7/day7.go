package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type bagType struct {
	qte  int
	name string
}

type bagContentType struct {
	bagsContained [4]bagType
	tested        bool // To avoid recursing multiple time
	containsShiny bool
}

func addBag(line string) (string, bagContentType) {
	// extract infos (and no, no regex, I don't like regex)
	lineLen := len(line) - 1 // Remove final .
	datas := strings.Split(line[0:lineLen], " contain ")
	var bagContent bagContentType

	name := datas[0]
	pos := strings.LastIndex(name, " ")
	name = name[0:pos] // remove bags

	// Test if empty
	if datas[1] != "no other" {

		var bag bagType
		bags := strings.Split(datas[1], ", ")
		cpt := 0
		for _, val := range bags {
			index := strings.Index(val, " ")
			bag.name = val[index+1:]
			pos := strings.LastIndex(bag.name, " ")
			bag.name = bag.name[0:pos] // remove "bag(s)"
			qte, _ := strconv.Atoi(val[0:index])
			bag.qte = qte
			bagContent.bagsContained[cpt] = bag
			cpt++
		}
	}
	return name, bagContent
}

func part1(bagList map[string]bagContentType, bagToTest string) bool {
	for _, bag := range bagList[bagToTest].bagsContained {
		// Test if empty
		if bag.qte == 0 {
			return false
		}
		// Test if shiny bag
		if bag.name == "shiny gold" {
			return true
		}
		test := part1(bagList, bag.name)
		if test {
			return true
		}
	}
	return false
}

func part2(bagList map[string]bagContentType, bagToTest string) int {

	total := 1 // The current bag must be counted
	if bagList[bagToTest].bagsContained[0].qte != 0 {
		for _, bag := range bagList[bagToTest].bagsContained {
			if bag.qte != 0 {
				total += bag.qte * part2(bagList, bag.name)
			}
		}
	}
	return total
}

func main() {

	bagList := make(map[string]bagContentType)

	file, err := os.Open("day7.data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		index, content := addBag(line)
		bagList[index] = content
	}

	totalP1 := 0

	for key := range bagList {
		if part1(bagList, key) {
			totalP1++
		}
	}

	fmt.Println("Part 1 : ", totalP1)
	fmt.Println("Part 2 : ", part2(bagList, "shiny gold")-1) // Minus 1 becaus we don't count the shiny gold bag

}
