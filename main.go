package main

import (
	"fmt"
	"math"
	"os"
	"sort"
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

// https://adventofcode.com/2024/day/1#part1
func partOne(lst1 []int, lst2 []int) {
	sort.Ints(lst1)
	sort.Ints(lst2)

	sum := 0
	for i, v := range lst1 {
		fmt.Printf("v=%d\n", v)
		sum += int(math.Abs(float64(v - lst2[i])))
	}
	fmt.Printf("answer: %d\n", sum)
}

// https://adventofcode.com/2024/day/1#part2
func partTwo(lst1 []int, lst2 []int) {
	// count the number of occuring in lst1
	dict1 := make(map[int]int)
	for _, v := range lst1 {
		dict1[v] = dict1[v] + 1
	}
	// count the number of occuring in lst2
	dict2 := make(map[int]int)
	for _, v := range lst2 {
		dict2[v] = dict2[v] + 1
	}

	// calculate similarity score
	var similarityScore int
	for k, v := range dict2 {
		fmt.Printf("k=%d, v=%v", k, v)
		if dict1[k] > 0 {
			fmt.Printf(" ---> found in lst1 (%d)!", dict1[k])
			similarityScore += k * v * dict1[k]
		}
		fmt.Printf("\n")
	}
	fmt.Printf("answer: %d\n", similarityScore)
}

func main() {
	inputStr := formatInput(readFile("numbers.txt"))

	// 1. format input
	var lst1, lst2 []int
	for _, v := range inputStr {
		pair := strings.Split(v, "   ")
		i1, err1 := strconv.Atoi(pair[0])
		i2, err2 := strconv.Atoi(pair[1])
		if err1 == nil && err2 == nil {
			lst1 = append(lst1, i1)
			lst2 = append(lst2, i2)
		} else {
			fmt.Print("Error", err1, err2)
		}
	}

	// start

	// Part 1
	partOne(lst1, lst2)

	// Part 2
	partTwo(lst1, lst2)
}
