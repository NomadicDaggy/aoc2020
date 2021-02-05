package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func passportHasRequiredFields(passport map[string]string) bool {
	requiredKeys := [7]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, key := range requiredKeys {
		_, ok := passport[key]
		if !ok {
			return false
		}
	}
	return true
}

func isAlphanumeric(s string) bool {
	match, _ := regexp.MatchString("^[0-9]+$", s)
	return match
}

func passportFieldsValid(passport map[string]string) bool {
	// Takes for granted that all required fields have been checked before
	if !isAlphanumeric(passport["byr"]) {
		return false
	}
	byr, _ := strconv.Atoi(passport["byr"])
	if byr < 1920 || byr > 2002 {
		return false
	}

	if !isAlphanumeric(passport["iyr"]) {
		return false
	}
	iyr, _ := strconv.Atoi(passport["iyr"])
	if iyr < 2010 || iyr > 2020 {
		return false
	}

	if !isAlphanumeric(passport["eyr"]) {
		return false
	}
	eyr, _ := strconv.Atoi(passport["eyr"])
	if eyr < 2020 || eyr > 2030 {
		return false
	}

	hgt := passport["hgt"]
	hgtVal, _ := strconv.Atoi(hgt[:len(hgt)-2])
	hgtUnit := hgt[len(hgt)-2:]
	if (hgtUnit == "cm" && (hgtVal < 150 || hgtVal > 193)) ||
		(hgtUnit == "in" && (hgtVal < 59 || hgtVal > 76)) {
		return false
	}
	if hgtUnit != "cm" && hgtUnit != "in" {
		return false
	}

	hcl := passport["hcl"]
	if hcl[0] != []byte("#")[0] {
		return false
	}
	if len(hcl) != 7 {
		return false
	}
	match, _ := regexp.MatchString("^[a-z0-9]*$", hcl[1:])
	if !match {
		return false
	}

	ecl := passport["ecl"]
	allowedEyeColors := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
	_, ok := allowedEyeColors[ecl]
	if !ok {
		return false
	}

	pid := passport["pid"]
	if len(pid) != 9 {
		return false
	}
	if !isAlphanumeric(pid) {
		return false
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
		if passportHasRequiredFields(passport) && passportFieldsValid(passport) {
			validPassportCount++
		}
	}

	fmt.Println(validPassportCount)
}
