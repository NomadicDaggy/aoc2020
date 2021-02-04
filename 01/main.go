package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func FindTwo(ints []int) int {
	for iIndex, i := range ints {
		for jIndex := iIndex + 1; jIndex < len(ints); jIndex++ {
			if j := ints[jIndex]; i+j == 2020 {
				return i * j
			}
		}
	}
	return 0
}

func FindThree(ints []int) int {
	for iIndex, i := range ints {
		for jIndex := iIndex + 1; jIndex < len(ints); jIndex++ {
			for kIndex := jIndex + 1; kIndex < len(ints); kIndex++ {
				j := ints[jIndex]
				k := ints[kIndex]
				if i+j+k == 2020 {
					return i * j * k
				}
			}
		}
	}
	return 0
}

func main() {
	fptr := flag.String("fpath", "input", "read input file")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	ints, err := ReadInts(f)
	fmt.Println(FindTwo(ints))
	fmt.Println(FindThree(ints))
}
