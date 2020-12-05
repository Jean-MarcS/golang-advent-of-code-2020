package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//May be should I used a map instead ?
type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func isValidPassportP1(currentPassport passport) bool {

	if currentPassport.byr == "" {
		return false
	}
	if currentPassport.iyr == "" {
		return false
	}
	if currentPassport.eyr == "" {
		return false
	}
	if currentPassport.hgt == "" {
		return false
	}
	if currentPassport.hcl == "" {
		return false
	}
	if currentPassport.ecl == "" {
		return false
	}
	if currentPassport.pid == "" {
		return false
	}
	return true
}

func isValidPassportP2(currentPassport passport) bool {

	if currentPassport.byr == "" {
		return false
	} else {
		byr, _ := strconv.Atoi(currentPassport.byr)
		if byr < 1920 || byr > 2002 {
			return false
		}
	}
	if currentPassport.iyr == "" {
		return false
	} else {
		iyr, _ := strconv.Atoi(currentPassport.iyr)
		if iyr < 2010 || iyr > 2020 {
			return false
		}
	}
	if currentPassport.eyr == "" {
		return false
	} else {
		eyr, _ := strconv.Atoi(currentPassport.eyr)
		if eyr < 2020 || eyr > 2030 {
			return false
		}
	}
	if currentPassport.hgt == "" {
		return false
	} else {
		pos := strings.Index(currentPassport.hgt, "in")
		if pos > 0 {
			height, _ := strconv.Atoi(currentPassport.hgt[:pos])
			if height < 59 || height > 76 {
				return false
			}
		} else {
			pos = strings.Index(currentPassport.hgt, "cm")
			if pos > 0 {
				height, _ := strconv.Atoi(currentPassport.hgt[:pos])
				if height < 150 || height > 193 {
					return false
				}
			} else {
				return false
			}
		}
	}
	if currentPassport.hcl == "" {
		return false
	} else {
		if len(currentPassport.hcl) == 7 {
			if currentPassport.hcl[0:1] == "#" {
			} else {
				for i := 1; i < 7; i++ {
					if !((currentPassport.hcl[i] >= '0' && currentPassport.hcl[i] <= '9') || (currentPassport.hcl[i] >= 'a' && currentPassport.hcl[i] <= 'f')) {
						return false
					}
				}
			}
		} else {
			return false
		}
	}

	if currentPassport.ecl == "" {
		return false
	} else {
		if currentPassport.ecl != "amb" && currentPassport.ecl != "blu" && currentPassport.ecl != "brn" && currentPassport.ecl != "gry" && currentPassport.ecl != "grn" && currentPassport.ecl != "hzl" && currentPassport.ecl != "oth" {
			return false
		}
	}
	if len(currentPassport.pid) != 9 {
		return false
	}
	return true
}

func main() {

	//var passports []passport

	currentPassport := passport{byr: "", iyr: "", eyr: "", hgt: "", hcl: "", ecl: "", pid: "", cid: ""}
	file, err := os.Open("day4.data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalValidPassportP1 := 0
	totalValidPassportP2 := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 1 {
			datas := strings.Split(line, " ")
			for _, val := range datas {
				info := strings.Split(val, ":")
				switch info[0] {
				case "byr":
					currentPassport.byr = info[1]
				case "iyr":
					currentPassport.iyr = info[1]
				case "eyr":
					currentPassport.eyr = info[1]
				case "hgt":
					currentPassport.hgt = info[1]
				case "hcl":
					currentPassport.hcl = info[1]
				case "ecl":
					currentPassport.ecl = info[1]
				case "pid":
					currentPassport.pid = info[1]
				case "cid":
					currentPassport.cid = info[1]
				}
			}
		} else {
			if isValidPassportP1(currentPassport) {
				totalValidPassportP1++
			}
			if isValidPassportP2(currentPassport) {
				totalValidPassportP2++
			}
			currentPassport = passport{byr: "", iyr: "", eyr: "", hgt: "", hcl: "", ecl: "", pid: "", cid: ""}
		}
	}

	fmt.Println("Part 1 : ", totalValidPassportP1)
	fmt.Println("Part 2 : ", totalValidPassportP2)

}
