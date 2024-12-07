package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var reports [][]int
	file, err := os.Open("inputs/2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineAsSlice := strings.Fields(line)
		reports = append(reports, toSliceOfInts(lineAsSlice))

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(PartTwo(reports))
}

func PartTwo(reports [][]int) int {
	var safeReports int
	for _, report := range reports {
		if getIsReportSafe(report, false) {
			safeReports++
		}

	}
	return safeReports
}

func getIsReportSafe(report []int, alreadyRemovedIndex bool) bool {
	isSafeReport := true
	isIncreasing := false

	for ndx, n := range report {
		if ndx == 0 {
			continue
		} else if ndx == 1 {
			isSafeReport = isDiffValid(n, report[ndx-1])
			isIncreasing = isReportTrendIncreasing(n, report[ndx-1])
		} else {
			isSafeReport = isDiffValid(n, report[ndx-1])

			currentTrendIsIncreasing := isReportTrendIncreasing(n, report[ndx-1])
			if currentTrendIsIncreasing != isIncreasing {
				isSafeReport = false
			}

		}

		if !isSafeReport {
			if !alreadyRemovedIndex {
				// part two
				for rdx, _ := range report {
					arrWithoutIndex := append([]int{}, report[:rdx]...)
					arrWithoutIndex = append(arrWithoutIndex, report[rdx+1:]...)
					isSafeReport = getIsReportSafe(arrWithoutIndex, true)
					if isSafeReport {
						break
					}
				}
			}
			break
		}
	}

	return isSafeReport
}

func toSliceOfInts(slice []string) []int {
	var res []int
	for _, s := range slice {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, n)
	}
	return res
}

func isDiffValid(currentN, prevN int) bool {
	diff := math.Abs(float64(currentN - prevN))
	if 1 <= diff && diff <= 3 {
		return true
	}
	return false
}

func isReportTrendIncreasing(currentN, prevN int) bool {
	if currentN > prevN {
		return true
	}
	return false
}

func reverseSlice[T any](arr []T) []T {
	reversed := []T{}
	for i := len(arr) - 1; i >= 0; i-- {
		reversed = append(reversed, arr[i])
	}
	return reversed
}
