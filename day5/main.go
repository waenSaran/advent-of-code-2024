package main

import (
	"fmt"
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

func topological_sort(rules [][]int, updates [][]int) [][]int {
	rules_map := make(map[int][]int)
	for _, rule := range rules {
		rules_map[rule[0]] = append(rules_map[rule[0]], rule[1])
	}
	new_updates := [][]int{}
	for _, update := range updates {
		for i := 0; i < len(update); i++ {
			for j := 0; j < len(update); j++ {
				if _, ok := rules_map[update[i]]; !ok {
					break
				}
				come_after_compared := i < j && slices.Contains(rules_map[update[i]], update[j])
				come_before_compared := i > j && slices.Contains(rules_map[update[j]], update[i])
				is_matched_rule := come_after_compared || come_before_compared
				if update[i] == update[j] {
					continue
				}
				if !is_matched_rule {
					// reorder
					update[i], update[j] = update[j], update[i]
				}
			}
		}
		new_updates = append(new_updates, update)
	}
	return new_updates
}

// https://adventofcode.com/2024/day/5#part2
func partTwo(rules [][]int, updates [][]int) {
	sum := 0
	wrong_updates := [][]int{}
	for _, update := range updates {
		if !is_follow_rules(rules, update) {
			// kept to new array
			wrong_updates = append(wrong_updates, update)
		}
	}
	// re-ordering update and sum mid nums
	new_updates := topological_sort(rules, wrong_updates)

	for _, update := range new_updates {
		// get middle nums
		mid_index := len(update) / 2
		sum += update[mid_index]
	}

	fmt.Printf("answer: %d\n", sum)
}

func main() {

	rules, updates := formatInput(readFile("input.txt"))

	// start
	partTwo(rules, updates)

}
