package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const height = 90

func nbAdjacentOccupied(array1 [height]string, l int, c int, width int, height int) int {
	total := 0
	if c > 0 && array1[l][c-1] == '#' {
		total++
	}
	if c < width-1 && array1[l][c+1] == '#' {
		total++
	}
	if l > 0 {
		if c > 0 && array1[l-1][c-1] == '#' {
			total++
		}
		if array1[l-1][c] == '#' {
			total++
		}
		if c < width-1 && array1[l-1][c+1] == '#' {
			total++
		}

	}
	if l < height-1 {
		if c > 0 && array1[l+1][c-1] == '#' {
			total++
		}
		if array1[l+1][c] == '#' {
			total++
		}
		if c < width-1 && array1[l+1][c+1] == '#' {
			total++
		}

	}
	return total
}

// Test all directions
func nbOccupied(array1 [height]string, line int, col int, width int, height int) int {
	total := 0
	// Up
	for l := line - 1; l >= 0; l-- {
		if array1[l][col] == 'L' {
			break
		}
		if array1[l][col] == '#' {
			total++
			break
		}
	}
	// Down
	for l := line + 1; l < height; l++ {
		if array1[l][col] == 'L' {
			break
		}
		if array1[l][col] == '#' {
			total++
			break
		}
	}
	// Left
	for c := col - 1; c >= 0; c-- {
		if array1[line][c] == 'L' {
			break
		}
		if array1[line][c] == '#' {
			total++
			break
		}
	}

	// Right
	for c := col + 1; c < width; c++ {
		if array1[line][c] == 'L' {
			break
		}
		if array1[line][c] == '#' {
			total++
			break
		}
	}

	// Up - Left
	for l, c := line-1, col-1; l >= 0 && c >= 0; l-- {
		if array1[l][c] == 'L' {
			break
		}
		if array1[l][c] == '#' {
			total++
			break
		}
		c--
	}

	// Up - Right
	for l, c := line-1, col+1; l >= 0 && c < width; l-- {
		if array1[l][c] == 'L' {
			break
		}
		if array1[l][c] == '#' {
			total++
			break
		}
		c++
	}

	// Down - Left
	for l, c := line+1, col-1; l < height && c >= 0; l++ {
		if array1[l][c] == 'L' {
			break
		}
		if array1[l][c] == '#' {
			total++
			break
		}
		c--
	}

	// Down - Right
	for l, c := line+1, col+1; l < height && c < width; l++ {
		if array1[l][c] == 'L' {
			break
		}
		if array1[l][c] == '#' {
			total++
			break
		}
		c++
	}

	return total
}

func part1(array1 [height]string) ([height]string, bool) {
	hasChange := false
	width := len(array1[0])

	array2 := array1
	for l := 0; l < height; l++ {
		for c := 0; c < width; c++ {
			switch array1[l][c] {
			case 'L':
				// exchange adjacent seat
				found := nbAdjacentOccupied(array1, l, c, width, height)

				if found == 0 { // No adjacent seat is occupied
					hasChange = true
					if c > 0 {
						exchange := array2[l]
						if c < width-1 {
							exchange = exchange[:c] + "#" + exchange[c+1:]
							array2[l] = exchange

						} else {
							exchange = exchange[:c] + "#" + exchange[c+1:]
							array2[l] = exchange
						}
					} else {
						exchange := array2[l]
						if c < width-1 {
							exchange = "#" + exchange[1:]
							array2[l] = exchange

						}
					}

				}
			case '#':
				found := nbAdjacentOccupied(array1, l, c, width, height)

				if found >= 4 { // four or more seat are occupied

					hasChange = true
					if c > 0 {
						exchange := array2[l]
						if c < width-1 {
							exchange = exchange[:c] + "L" + exchange[c+1:]
							array2[l] = exchange

						} else {
							exchange = exchange[:c] + "L" + exchange[c+1:]
							array2[l] = exchange
						}
					} else {
						exchange := array2[l]
						if c < width-1 {
							exchange = "L" + exchange[1:]
							array2[l] = exchange

						}
					}

				}
			}
		}

	}
	return array2, hasChange
}

func part2(array1 [height]string) ([height]string, bool) {
	hasChange := false
	width := len(array1[0])
	//fmt.Println("Width : ", width)

	array2 := array1
	for l := 0; l < height; l++ {
		for c := 0; c < width; c++ {
			switch array1[l][c] {
			case 'L':
				// exchange adjacent seat
				found := nbOccupied(array1, l, c, width, height)

				if found == 0 { // No visible seat is occupied
					hasChange = true
					if c > 0 {
						exchange := array2[l]
						if c < width-1 {
							exchange = exchange[:c] + "#" + exchange[c+1:]
							array2[l] = exchange

						} else {
							exchange = exchange[:c] + "#" + exchange[c+1:]
							array2[l] = exchange
						}
					} else {
						exchange := array2[l]
						if c < width-1 {
							exchange = "#" + exchange[1:]
							array2[l] = exchange

						}
					}

				}
			case '#':
				found := nbOccupied(array1, l, c, width, height)

				if found >= 5 { // five or more seat are occupied
					hasChange = true
					if c > 0 {
						exchange := array2[l]
						if c < width-1 {
							exchange = exchange[:c] + "L" + exchange[c+1:]
							array2[l] = exchange

						} else {
							exchange = exchange[:c] + "L" + exchange[c+1:]
							array2[l] = exchange
						}
					} else {
						exchange := array2[l]
						if c < width-1 {
							exchange = "L" + exchange[1:]
							array2[l] = exchange

						}
					}

				}
			}

		}
	}
	return array2, hasChange
}

//  Count occupied seats
func nbOccupiedSeats(array1 [height]string) int {
	ocSeats := 0
	width := len(array1[0])

	for l := 0; l < height; l++ {
		for c := 0; c < width; c++ {
			if array1[l][c] == '#' {
				ocSeats++
			}
		}
	}
	return ocSeats
}

func main() {

	var arraySrc [height]string

	file, err := os.Open("day11.data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cpt := 0

	for scanner.Scan() {
		arraySrc[cpt] = scanner.Text()
		cpt++
	}

	// Part 1
	array1 := arraySrc
	array2 := arraySrc
	hasChange := true
	for hasChange {
		array2, hasChange = part1(array1)
		if hasChange {
			array1 = array2
		}
	}

	fmt.Println("Part 1 : ", nbOccupiedSeats(array1))

	// Part 2
	array1 = arraySrc
	array2 = arraySrc
	hasChange = true
	for hasChange {
		array2, hasChange = part2(array1)
		if hasChange {
			array1 = array2
		}
	}

	fmt.Println("Part 2 : ", nbOccupiedSeats(array1))

}
