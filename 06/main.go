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

func isRuneInString(slice string, charRune rune) bool {
	for _, item := range slice {
		if item == charRune {
			return true
		}
	}
	return false
}

func main() {
	lines, _ := readLines("input")
	yesAnswers := make(map[rune]bool)
	cumulativeYesAnswers := 0
	lineCount := len(lines)
	for i, line := range lines {
		for _, char := range line {
			// if char is already in map, this does nothing
			// if line is empty this also does nothing
			yesAnswers[char] = true
		}
		if line == "" || i == lineCount-1 {
			cumulativeYesAnswers += len(yesAnswers)
			yesAnswers = make(map[rune]bool)
			continue
		}
	}
	fmt.Println(cumulativeYesAnswers)

	cumulativeGroupwideYesCount := 0
	nextLineNewGroup := false
	// TODO: in retrospect, a map[string]bool would be a much better fit than
	// a rune array here. Easier to remove specific elements later on (just set to false)
	groupwideYesChars := []rune(lines[0])
	for i, line := range lines[1:] {
		if line == "" || i == lineCount-2 {
			cumulativeGroupwideYesCount += len(groupwideYesChars)
			nextLineNewGroup = true
			continue
		}
		if nextLineNewGroup {
			groupwideYesChars = []rune(line)
			nextLineNewGroup = false
		} else {
			for _, char := range groupwideYesChars {
				found := isRuneInString(line, char)
				if !found && len(groupwideYesChars) <= 1 {
					groupwideYesChars = []rune{}
					break
				} else if !found {
					// to delete a rune, we first have to find it
					// as we might have moved it while deleting
					// other runes
					indexToDelete := -1
					for j, runeToDelete := range groupwideYesChars {
						if runeToDelete == char {
							indexToDelete = j
						}
					}
					lastIndex := len(groupwideYesChars) - 1
					groupwideYesChars[indexToDelete] = groupwideYesChars[lastIndex]
					groupwideYesChars = groupwideYesChars[:lastIndex]
				}
			}
		}
	}
	fmt.Println(cumulativeGroupwideYesCount)
}
