package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var leftVals []int
	var rightVals []int

	file, err := os.Open("inputs/1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sections := strings.Fields(line)
		if len(sections) != 2 {
			continue
		}

		leftN, err := strconv.Atoi(sections[0])
		if err != nil {
			log.Fatal(err)
		}
		rightN, err := strconv.Atoi(sections[1])
		if err != nil {
			log.Fatal(err)
		}

		rightVals = append(rightVals, rightN)
		leftVals = append(leftVals, leftN)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(partTwo(leftVals, rightVals))
}

func partTwo(leftVals, rightVals []int) int {
	res := 0
	rightValCountHash := map[int]int{} // number:count

	for _, n := range rightVals {
		rightValCountHash[n]++
	}

	for _, n := range leftVals {
		res += n * rightValCountHash[n]
	}

	return res
}

func partOne(leftVals, rightVals []int) int {
	var answer int
	sort.Ints(rightVals)
	for i := 0; i < len(leftVals); i++ {
		difference := diff(leftVals[i], rightVals[i])
		answer += difference
	}
	return answer

}

func diff(n1, n2 int) int {
	if n1 > n2 {
		return n1 - n2
	}
	return n2 - n1
}
