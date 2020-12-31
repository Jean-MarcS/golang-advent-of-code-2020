package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type memory struct {
	address int
	value   int
}

func applyMask(mask string, val int) int {
	for i := 0; i < 36; i++ {
		c := 1 << i
		currentBit := (val & c) >> i
		switch mask[35-i] {
		case 48: // 0
			if currentBit == 1 {
				val -= c
			}
		case 49: //1
			if currentBit == 0 {
				val += c
			}
		}
	}

	return val
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func getMaskAddresses(mask string, address int) []int {
	var addressList []int
	maskOut := mask
	nbX := 0
	for i := 0; i < 36; i++ {
		c := 1 << i
		currentBit := (address & c) >> i
		switch mask[35-i] {
		case 48: // 0
			char := "0"
			if currentBit == 1 {
				char = "1"
			}
			maskOut = maskOut[0:35-i] + char + maskOut[35-i+1:]
		case 49: //1
			maskOut = maskOut[0:35-i] + "1" + maskOut[35-i+1:]
		default:
			maskOut = maskOut[0:35-i] + "X" + maskOut[35-i+1:]
			nbX++
		}
	}
	// Loop possibilities
	var b int
	comb := 1 << nbX
	for b = 0; b < comb; b++ {
		st := fmt.Sprintf("%016b", b)
		newAddress := maskOut
		// Find all X
		pos := strings.LastIndex(newAddress, "X")
		cpt := len(st) - 1
		for pos >= 0 {
			end := newAddress[pos+1:]
			newAddress = newAddress[0:pos] + st[cpt:cpt+1] + end
			cpt--
			pos = strings.LastIndex(newAddress, "X")

		}
		// Get int value
		na, _ := strconv.ParseInt(newAddress, 2, 64)
		addressList = append(addressList, int(na))

	}
	return addressList
}

func main() {

	mem := []memory{}
	memP2 := []memory{}

	file, err := os.Open("day14.data")
	if err != nil {
		log.Fatal(err)
	}

	// Part 1
	scanner := bufio.NewScanner(file)
	currentMask := ""

	for scanner.Scan() {
		line := scanner.Text()
		if line[0:4] == "mask" {
			currentMask = line[7:]
		} else {
			pos := strings.Index(line, "]")
			address, _ := strconv.Atoi(line[4:pos])
			value, _ := strconv.Atoi(line[pos+4:])
			valueP1 := applyMask(currentMask, value)

			// Part 1
			found := false
			goodKey := 0
			for key, val := range mem {
				if val.address == address {
					found = true
					goodKey = key
					break
				}
			}
			if found {
				mem[goodKey].value = valueP1
			} else {
				mem = append(mem, memory{address: address, value: valueP1})
			}

			// Part2
			adressList := getMaskAddresses(currentMask, address)

			for _, memAdress := range adressList {
				found := false
				goodKey := 0
				for key, val := range memP2 {
					if val.address == memAdress {
						found = true
						goodKey = key
						break
					}
				}
				if found {
					memP2[goodKey].value = value
				} else {
					memP2 = append(memP2, memory{address: memAdress, value: value})
				}
			}

		}
	}

	totalP1 := 0
	for _, val := range mem {
		totalP1 += val.value
	}

	fmt.Println("Part 1 : ", totalP1)

	totalP2 := 0
	for _, val := range memP2 {
		totalP2 += val.value
	}

	fmt.Println("Part 2 : ", totalP2)
}
