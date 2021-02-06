package main

import (
	"bufio"
	"fmt"
	"os"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func binarySeek(lowSymb rune, highSymb rune, lowThresh int, highThresh int, boardingPass string) int {
	for _, symbol := range boardingPass {
		diff := highThresh - lowThresh
		if symbol == lowSymb {
			if diff == 1 {
				return lowThresh
			}
			highThresh -= (diff + 1) / 2 // diff+1 to cheat integer division
		} else if symbol == highSymb {
			if diff == 1 {
				return highThresh
			}
			lowThresh += (diff + 1) / 2
		} else {
			return -1
		}
		//fmt.Println(lowThresh, highThresh)
	}
	return highThresh
}

func getSeatID(boardingPass string) int {
	row := binarySeek('F', 'B', 0, 127, boardingPass[:7])
	col := binarySeek('L', 'R', 0, 7, boardingPass[7:])
	//fmt.Println(boardingPass, boardingPass[:7], boardingPass[7:], row, col)
	return row*8 + col
}

func main() {
	lines, _ := readLines("input")
	highestID := 0
	for _, boardingPass := range lines {
		//fmt.Println("")
		id := getSeatID(boardingPass)
		//fmt.Println(id)
		if id > highestID {
			highestID = id
		}
	}
	fmt.Println(highestID)
}
