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

func countTrees(right int, down int, treeMap []string) int {
	xLastIndex := len(treeMap[0]) - 1
	yLastIndex := len(treeMap) - 1
	treeCount := 0
	x, y := 0, 0
	for {
		x += right
		y += down
		if x > xLastIndex {
			x -= xLastIndex + 1
		}
		if y > yLastIndex {
			break
		}
		if treeMap[y][x] == []byte("#")[0] {
			treeCount++
		}
	}
	return treeCount
}

func main() {
	lines, _ := readLines("input")
	fmt.Println(countTrees(3, 1, lines))
}
