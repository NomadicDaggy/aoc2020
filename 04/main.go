package main

import (
	"bufio"
	"fmt"
	"os"
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

func isPassportValid(passport map[string]string) bool {
	requiredKeys := [7]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, key := range requiredKeys {
		_, ok := passport[key]
		if !ok {
			return false
		}
	}
	return true
}

func main() {
	lines, _ := readLines("input")
	//var passports []map[string]string
	var passportStrings []string
	activePassportIdx := 0
	nextLineNewPassport := true
	for _, line := range lines {
		if line == "" {
			nextLineNewPassport = true
			activePassportIdx++
			continue
		}
		if nextLineNewPassport {
			passportStrings = append(passportStrings, line)
		} else {
			unfinishedPassport := passportStrings[activePassportIdx]
			passportStrings[activePassportIdx] = unfinishedPassport + " " + line
		}
		nextLineNewPassport = false
	}

	var passports []map[string]string
	for _, passportString := range passportStrings {
		fieldStrings := strings.Split(passportString, " ")
		passportMap := make(map[string]string)
		for _, fieldString := range fieldStrings {
			s := strings.Split(fieldString, ":")
			passportMap[s[0]] = s[1]
		}
		passports = append(passports, passportMap)
	}

	validPassportCount := 0
	for _, passport := range passports {
		if isPassportValid(passport) {
			validPassportCount++
		}
	}

	fmt.Println(validPassportCount)
}
