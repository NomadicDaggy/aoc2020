package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/graph/simple"
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
	// fmt.Println(bagContentsMap)
	// fmt.Println("")

	graph := simple.NewDirectedGraph()
	bagColorToNodeID := make(map[string]int64)
	for parentBagColor, innerMap := range bagContentsMap {
		colorID, exists := bagColorToNodeID[parentBagColor]
		parentNode := graph.NewNode()
		if !exists {
			bagColorToNodeID[parentBagColor] = parentNode.ID()
			graph.AddNode(parentNode)
		} else {
			parentNode = graph.Node(colorID)
		}

		for childBagColor, numContains := range innerMap {
			if numContains > 0 {
				childNode := graph.NewNode()
				childColorID, childExists := bagColorToNodeID[childBagColor]
				if !childExists {
					bagColorToNodeID[childBagColor] = childNode.ID()
					// No need to explicitly add child node as setting the edge
					// will add it automatically
				} else {
					childNode = graph.Node(childColorID)
				}
				edge := graph.NewEdge(parentNode, childNode)
				graph.SetEdge(edge)
			}
		}
	}
	shinyGoldID := bagColorToNodeID["shiny gold"]
	leadsToShinyGold := make(map[string]bool)
	for color, nodeID := range bagColorToNodeID {
		if nodeID == shinyGoldID {
			continue
		}
		if graph.HasEdgeFromTo(nodeID, shinyGoldID) {
			leadsToShinyGold[color] = true
		}
	}
	fmt.Println(leadsToShinyGold)

	// TODO: https://pkg.go.dev/gonum.org/v1/gonum@v0.8.2/graph/path
}
