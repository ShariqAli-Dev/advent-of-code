package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const DIAL_START = 50

func main() {
	inputFileBytes, err := os.ReadFile("inputs/1.txt")
	if err != nil {
		log.Fatalf("err opening file: %v", err)
	}
	inputFile := string(inputFileBytes)
	sequenceSlice := strings.Split(inputFile, "\n")
	timesPassedZero := Answer(sequenceSlice)

	fmt.Printf("times passed zero: %d\n", timesPassedZero)
}

func Answer(sequenceSlice []string) (timesPassedZero int) {
	dial := DIAL_START

	for _, sequence := range sequenceSlice {
		if len(sequence) == 0 {
			continue
		}

		turnDirection := sequence[:1]
		turns, err := strconv.Atoi(sequence[1:len(sequence)])
		if err != nil {
			log.Fatalf("err extracting sequence turns: %v", err)
		}

		if turnDirection == "L" {
			dial = (dial - turns%100 + 100) % 100
		} else if turnDirection == "R" {
			dial = (dial + turns) % 100
		} else {
			log.Fatalf("err invalid rotation direction: %s", turnDirection)
		}
		if dial == 0 {
			timesPassedZero++
		}

		fmt.Printf("dial is rotated %s%d to point at %d\n", turnDirection, turns, dial)
	}
	return timesPassedZero
}
