package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func IsPasswordValid(s string) bool {
	parts := strings.Split(s, ": ")
	password := parts[1]
	prefixParts := strings.Split(parts[0], " ")
	char := []rune(prefixParts[1])[0]
	boundParts := strings.Split(prefixParts[0], "-")
	lowerBound, _ := strconv.Atoi(boundParts[0])
	upperBound, _ := strconv.Atoi(boundParts[1])
	charCounter := 0
	for _, c := range password {
		if c == char {
			charCounter++
		}
		if charCounter > upperBound {
			return false
		}
	}
	return charCounter >= lowerBound
}

func IsNewPasswordValid(s string) bool {
	parts := strings.Split(s, ": ")
	password := parts[1]
	prefixParts := strings.Split(parts[0], " ")
	char := []byte(prefixParts[1])[0]
	indexParts := strings.Split(prefixParts[0], "-")
	lowerIndex, _ := strconv.Atoi(indexParts[0])
	upperIndex, _ := strconv.Atoi(indexParts[1])

	// for the requirement to hold, one char has to match and must not
	if (password[lowerIndex-1] == char) != (password[upperIndex-1] == char) {
		return true
	}
	return false
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

	reader := bufio.NewReader(f)
	correctPasswords := 0
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if IsNewPasswordValid(string(line)) {
			correctPasswords++
		}
	}
	fmt.Println(correctPasswords)
}
