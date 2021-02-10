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
	accBeforeInstructionAt := make(map[int]int)
	acc := 0
	i := 0
	var instructionOrder []int // needed for 08b
	// determine failure point (08a)
	for {
		instructionOrder = append(instructionOrder, i)
		if _, ok := accBeforeInstructionAt[i]; ok {
			fmt.Println(acc, " :failure point")
			accBeforeInstructionAt[i] = acc
			break
		}
		accBeforeInstructionAt[i] = acc
		i, acc = processInstruction(lines[i], acc, i)
	}
	// determine broken instruction (08b)
	// holds line indexes from which there is no path out of the boot code
	// blacklist := make(map[int]bool)

	// Search backwards through original path
	endIndex := len(lines)
	for j := len(instructionOrder) - 1; j > 0; j-- {
		// Change first changeable instruction and look for ending
		lineIndex := instructionOrder[j]
		instructionString := lines[lineIndex]
		switch instructionString[:3] {
		case "nop":
			instructionString = "jmp" + instructionString[3:]
		case "jmp":
			instructionString = "nop" + instructionString[3:]
		default:
			continue
		}

		// run with swapped instruction
		i = lineIndex
		acc, ok := accBeforeInstructionAt[lineIndex]
		if !ok {
			panic("malformed history")
		}
		i, acc = processInstruction(instructionString, acc, i)

		// continue until end found or infinite loop
		linesVisited := make(map[int]bool)
		for {
			if i == endIndex {
				fmt.Println(acc)
				return
			} else if _, ok := linesVisited[i]; ok {
				break
			}
			linesVisited[i] = true
			i, acc = processInstruction(lines[i], acc, i)
		}
	}
}
