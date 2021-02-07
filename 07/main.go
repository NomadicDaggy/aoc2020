package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func parseLines(lines []string) map[string](map[string]int) {
	outerMap := make(map[string](map[string]int))
	for _, line := range lines {
		line := strings.ReplaceAll(line, "bags", "")
		line = strings.ReplaceAll(line, "bag", "")
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, ".", "")
		line = strings.ReplaceAll(line, "  ", " ")
		line = strings.ReplaceAll(line, " no other", " 0 no other")

		split := strings.Split(line, " contain ")
		outerKey := split[0]                     // dark olive
		outerVal := strings.Split(split[1], " ") // [3 faded blue  4 dotted black ]

		innerMap := make(map[string]int)
		for i := 3; i < len(outerVal); i += 3 {
			bagName := outerVal[i-2] + " " + outerVal[i-1]
			bagCount, _ := strconv.Atoi(outerVal[i-3])
			// fmt.Println(bagName, "--", bagCount)
			// fmt.Println("")
			innerMap[bagName] = bagCount
		}
		outerMap[outerKey] = innerMap
	}
	return outerMap
}

func main() {
	lines, _ := readLines("example")
	bagContentsMap := parseLines(lines)

	fmt.Println(bagContentsMap)
}
