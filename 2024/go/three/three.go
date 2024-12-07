package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var lines []string
	file, err := os.Open("../inputs/3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(PartOne(lines))
}

func PartOne(lines []string) int {
	var matches [][]string
	var matchString []string
	res := 0
	for _, line := range lines {
		pattern := `mul\((\d+),(\d+)\)`
		reg := regexp.MustCompile(pattern)
		matches = reg.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			matchString = append(matchString, match[0])
			num1, err := strconv.Atoi(match[1])
			if err != nil {
				log.Fatal(err)
			}
			num2, err := strconv.Atoi(match[2])
			if err != nil {
				log.Fatal(err)
			}
			res += num1 * num2

		}
	}
	return res
}
