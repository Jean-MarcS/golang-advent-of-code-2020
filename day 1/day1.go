package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day1.data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cpt := 0
	var lines [200]int

	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		lines[cpt] = val
		cpt++
	}

	cpt = 0
	finish := false
	for cpt < 200 && !finish {
		cptSum := cpt + 1
		for cptSum < 200 && !finish {
			if lines[cpt]+lines[cptSum] == 2020 {
				fmt.Println("Part 1 : ", lines[cpt]*lines[cptSum])
				finish = true
			}
			cptSum++
		}
		cpt++
	}
	cpt = 0
	finish = false
	for cpt < 200 && !finish {
		cptSum := cpt + 1
		for cptSum < 200 && !finish {
			cptSum3 := cptSum + 1
			for cptSum3 < 200 && !finish {
				if lines[cpt]+lines[cptSum]+lines[cptSum3] == 2020 {
					fmt.Println("Part 2 : ", lines[cpt]*lines[cptSum]*lines[cptSum3])
					finish = true
				}
				cptSum3++
			}
			cptSum++
		}
		cpt++
	}

}
