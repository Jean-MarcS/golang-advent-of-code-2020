package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func travel(stepX int, stepY int, width int, height int, field []string) int {
	currentX := 0
	currentY := 0
	nbTrees := 0
	for currentY < height-stepY {
		currentX = (currentX + stepX) % width
		currentY = currentY + stepY
		if field[currentY][currentX:currentX+1] == "#" {
			nbTrees++
		}
	}

	return nbTrees
}

func main() {

	totalSize := 323
	field := make([]string, totalSize+1)

	file, err := os.Open("day3.data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cpt := 0
	for scanner.Scan() {
		field[cpt] = scanner.Text()
		cpt++
	}

	nbTrees := travel(3, 1, len(field[0]), totalSize, field)
	fmt.Println("Part 1 : ", nbTrees)

	nbTrees = nbTrees * travel(1, 1, len(field[0]), totalSize, field)
	nbTrees = nbTrees * travel(5, 1, len(field[0]), totalSize, field)
	nbTrees = nbTrees * travel(7, 1, len(field[0]), totalSize, field)
	nbTrees = nbTrees * travel(1, 2, len(field[0]), totalSize, field)

	fmt.Println("Part 2 : ", nbTrees)

}
