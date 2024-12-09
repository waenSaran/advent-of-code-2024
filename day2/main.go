package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readFile(fileName string) string {
	b, err := os.ReadFile(fileName) // return in bytes
	if err != nil {
		fmt.Print(err)
		return ""
	}
	str := string(b)
	return str
}

func formatInput(s string) []string {
	return strings.Split(s, "\n")
}

/*
Unsafe report conditions
1. has equal value
2. has both increasing and decreasing
3. all increase but diff is greater than 3
4. all decrease but diff is greater than 3

Safe report condition
1. all decreasing and its diff are at least 1 and at most 3
2. all increasing and its diff are at least 1 and at most 3
*/
func isReportSafe(report []string) bool {
	fmt.Print("\nreport: ", report)
	isIncreasing := false
	for i, v := range report {
		if i == len(report)-1 {
			break
		}
		curr, err_c := strconv.Atoi(v)
		next, err_n := strconv.Atoi(report[i+1])
		if err_c != nil && err_n != nil {
			fmt.Print(err_c, err_n)
			return false
		}
		diff := curr - next
		abs_diff := math.Abs(float64(diff))
		if abs_diff > 3 || abs_diff < 1 {
			fmt.Print(" ---> BOUNDARY UNSAFE!!")
			return false
		}

		// Set increasing / decreasing flag
		if i == 0 && diff > 0 {
			isIncreasing = true
		}

		// check flag conflict
		isConflict := (isIncreasing && diff < 0) || (!isIncreasing && diff > 0)
		if i != 0 && isConflict {
			fmt.Print(" ---> CONFLICT UNSAFE!!")
			return false
		}
	}
	return true
}

func deleteElement(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

// Return does this report can safe by remove just 1 level
func canSolveUnSafeReport(report []string) bool {
	// remove each level and prove it is safe
	fmt.Print("\n-----try to solve this report----\n")
	for i := 0; i < len(report); i++ {
		tmp_report := slices.Clone(report)
		tmp_report = deleteElement(tmp_report, i)
		if isReportSafe(tmp_report) {
			fmt.Print("\nDone solve report with SAFE!\n")
			return true
		}
	}
	fmt.Print("\nCannot solve report\n")
	return false
}

// https://adventofcode.com/2024/day/2
func partOne(reports []string) {
	count_safe_report := 0
	for _, r := range reports {
		report := strings.Split(r, " ")
		isSafe := isReportSafe(report)
		if isSafe || (!isSafe && canSolveUnSafeReport(report)) {
			count_safe_report += 1
		}
		fmt.Print("\n")
	}
	fmt.Printf("answer: %d\n", count_safe_report)
}

func main() {
	// 1. formart input
	reports := formatInput(readFile("input.txt"))

	// start
	partOne(reports)

}
