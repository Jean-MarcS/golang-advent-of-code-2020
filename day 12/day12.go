package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type instruction struct {
	action byte
	value  int
}

const size = 788

func part1(lines [size]instruction) int {
	facing := 0
	x := 0
	y := 0

	for cpt := 0; cpt < size; cpt++ {
		switch lines[cpt].action {
		case 'N':
			y -= lines[cpt].value
		case 'S':
			y += lines[cpt].value
		case 'E':
			x += lines[cpt].value
		case 'W':
			x -= lines[cpt].value
		case 'L':
			facing = (facing + lines[cpt].value) % 360
		case 'R':
			facing = (facing + 360 - lines[cpt].value) % 360
		case 'F':
			switch facing {
			case 0:
				x += lines[cpt].value
			case 90:
				y -= lines[cpt].value
			case 180:
				x -= lines[cpt].value
			case 270:
				y += lines[cpt].value
			}
		}
	}

	return x + y
}

func part2(lines [size]instruction) int {

	x := 0
	y := 0

	wpx := 10
	wpy := -1

	for cpt := 0; cpt < size; cpt++ {
		switch lines[cpt].action {
		case 'N':
			wpy -= lines[cpt].value
		case 'S':
			wpy += lines[cpt].value
		case 'E':
			wpx += lines[cpt].value
		case 'W':
			wpx -= lines[cpt].value
		case 'L': // Inverted
			switch lines[cpt].value {
			case 90:
				dummy := -wpx
				wpx = wpy
				wpy = dummy
			case 180:
				wpx = -wpx
				wpy = -wpy
			case 270:
				dummy := -wpy
				wpy = wpx
				wpx = dummy
			}
		case 'R':
			switch lines[cpt].value {
			case 270:
				dummy := -wpx
				wpx = wpy
				wpy = dummy
			case 180:
				wpx = -wpx
				wpy = -wpy
			case 90:
				dummy := -wpy
				wpy = wpx
				wpx = dummy
			}
		case 'F':
			x += wpx * lines[cpt].value
			y += wpy * lines[cpt].value
		}

	}

	return int(math.Abs(float64(x + y)))
}

func main() {
	file, err := os.Open("day12.data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines [size]instruction

	cpt := 0
	for scanner.Scan() {
		line := scanner.Text()
		lines[cpt].action = line[0]
		lines[cpt].value, _ = strconv.Atoi(line[1:])
		cpt++
	}

	fmt.Println("Part 1 :", part1(lines))
	fmt.Println("Part 2 :", part2(lines))

}
