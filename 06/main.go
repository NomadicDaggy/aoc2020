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

func main() {
	lines, _ := readLines("input")
	yesAnswers := make(map[rune]bool)
	cumulativeYesAnswers := 0
	lineCount := len(lines)
	for i, line := range lines {
		// fmt.Println(line)
		for _, char := range line {
			// if char is already in map, this does nothing
			// if line is empty this also does nothing
			yesAnswers[char] = true
		}
		if line == "" || i == lineCount-1 {
			cumulativeYesAnswers += len(yesAnswers)
			yesAnswers = make(map[rune]bool)
			// fmt.Println(cumulativeYesAnswers)
			continue
		}
	}
	fmt.Println(cumulativeYesAnswers)
}
