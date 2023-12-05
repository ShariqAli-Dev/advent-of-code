package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const asciiZero = 48
const asciiNine = 57

var numberWords = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	file, err := os.Open("../inputs/one.txt")
	if err != nil {
		log.Fatal("could not open the file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var calibrationSum uint
	for scanner.Scan() {
		line := scanner.Text()
		var leftNumber string
		var rightNumber string
		for i := 0; i < len(line); i++ {
			lineAscii := line[i]
			if isNumericAscii(lineAscii) {
				leftNumber = string(lineAscii)
				break
			}
			// 2 star
			value, exists := numberWords[line[0:i]]
			if exists {
				leftNumber = value
				break
			}
		}
		for i := len(line); i > 0; i-- {
			lineAscii := line[i-1]
			if isNumericAscii(lineAscii) {
				rightNumber = string(lineAscii)
				break
			}
			// 2 star
			fmt.Println(line[i:])
			value, exists := numberWords[line[i:]]
			if exists {
				rightNumber = value
				break
			}
		}
		calibration, err := strconv.Atoi(leftNumber + rightNumber)
		if err != nil {
			log.Fatal(err)
		}
		calibrationSum += uint(calibration)
	}
	fmt.Println("calibration:", calibrationSum)
}
func isNumericAscii(ascii byte) bool {
	return ascii >= asciiZero && ascii <= asciiNine
}
