package main

import (
	"fmt"
	"os"
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

func formatInput(longtxt string) ([][]int, [][]int) {
	sec := strings.Split(longtxt, "\n\n")
	sec1_tmp := strings.Split(sec[0], "\n")
	sec2_tmp := strings.Split(sec[1], "\n")

	// format sec1
	sec1 := [][]int{}
	for _, rule_str := range sec1_tmp {
		rule_tmp := strings.Split(rule_str, "|")
		nums := []int{}
		for _, s := range rule_tmp {
			num, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}
		sec1 = append(sec1, nums)
	}

	// format sec2
	sec2 := [][]int{}
	for _, line_str := range sec2_tmp {
		new_line := []int{}
		line := strings.Split(line_str, ",")
		for _, str_num := range line {
			num, err := strconv.Atoi(string(str_num))
			if err != nil {
				panic(err)
			}
			new_line = append(new_line, num)
		}
		sec2 = append(sec2, new_line)
	}
	return sec1, sec2
}

func is_follow_rules(rules [][]int, update_pages []int) bool {
	for li := 0; li < len(update_pages); li++ {
		for ri := li + 1; ri < len(update_pages); ri++ {
			lpage := update_pages[li]
			rpage := update_pages[ri]
			// check rules
			match_rule := false
			for _, rule := range rules {
				if rule[0] == lpage && rule[1] == rpage {
					match_rule = true
					break
				}
			}
			if !match_rule {
				return false
			}
		}
	}
	return true
}

// https://adventofcode.com/2024/day/5
func partOne(rules [][]int, updates [][]int) {
	sum := 0
	for _, update := range updates {
		if is_follow_rules(rules, update) {
			// get middle nums
			mid_index := len(update) / 2
			sum += update[mid_index]
		}
	}
	fmt.Printf("answer: %d\n", sum)
}

func main() {

	rules, updates := formatInput(readFile("input.txt"))

	// start
	partOne(rules, updates)

}
