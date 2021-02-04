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
	for iIndex, i := range ints {
		for jIndex := iIndex; jIndex < len(ints); jIndex++ {
			if j := ints[jIndex]; i+j == 2020 {
				fmt.Println(i * j)
			}
		}
	}
}
