package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func stringToInstruction(line string) (string, int) {
	valInt, _ := strconv.Atoi(line[5:])
	if line[4] == []byte("-")[0] {
		valInt *= -1
	}
	return line[:3], valInt
}

// Processes an instruction, returns next instruction's index and current accumulator
func processInstruction(
	instructionString string,
	accumulator int,
	currentIndex int,
) (int, int) {
	instruction, argument := stringToInstruction(instructionString)
	switch instruction {
	case "nop":
		return currentIndex + 1, accumulator
	case "acc":
		return currentIndex + 1, accumulator + argument
	case "jmp":
		return currentIndex + argument, accumulator
	default:
		panic("unexpected instruction")
	}
}

func main() {
	lines, _ := readLines("input")
	visitedLines := make(map[int]bool)
	acc := 0
	i := 0
	for {
		if lineAlreadyVisited, _ := visitedLines[i]; lineAlreadyVisited {
			fmt.Println(acc)
			return
		}
		visitedLines[i] = true
		i, acc = processInstruction(lines[i], acc, i)
	}
}
