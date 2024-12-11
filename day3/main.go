package main

import (
	"fmt"
	"os"
	"regexp"
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

func findAllMul(txt string) [][]byte {
	content := []byte(txt)
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	return re.FindAll(content, -1)
}

func getNumbersInMul(txt []byte) []int {
	re := regexp.MustCompile(`\d+`)
	var numbers []int
	for _, v := range re.FindAll(txt, -1) {
		b2i, err := strconv.Atoi(string(v))
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, b2i)
	}
	return numbers
}

// get mul(), do() and don't()
func findAllFunc(txt string) [][]byte {
	content := []byte(txt)
	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	return re.FindAll(content, -1)
}

// https://adventofcode.com/2024/day/3
func partOne(longtxt string) {
	allMulString := findAllMul(longtxt)
	sum := 0
	for _, v := range allMulString {
		numbers := getNumbersInMul(v)
		sum += int(numbers[0]) * int(numbers[1])
	}
	fmt.Printf("sum: %d\n", sum)
}

// https://adventofcode.com/2024/day/3#part2
func partTwo(longtxt string) {
	allExecFunc := findAllFunc(longtxt)
	sum := 0
	do := true
	for _, v := range allExecFunc {
		if string(v) == "don't()" {
			do = false
		}
		if string(v) == "do()" {
			do = true
			continue
		}
		if !do {
			continue
		}
		numbers := getNumbersInMul(v)
		sum += int(numbers[0]) * int(numbers[1])
	}
	fmt.Printf("sum: %d\n", sum)
}

func formatInput(longtxt string) []string {
	return strings.Split(longtxt, "\n")
}

func main() {

	long_text := readFile("input.txt")

	// start
	partTwo(long_text)

}
