package main

import (
	"fmt"
	"os"
	"strings"
)

var STRING = "XMAS"
var LEN_STRING = len(STRING)
var DIRECTION = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

func is_out_of_range(x int, y int, len_row int, len_col int) bool {
	return x < 0 || y < 0 || x >= len_row || y >= len_col
}

func partOne(table [][]string) {
	count := 0
	for row_i, row := range table {
		for col_i, v := range row {
			// focus only X
			if v != "X" {
				continue
			}
			// find a valid direction
			for _, dir := range DIRECTION {
				found := false
				for i, s := range STRING {
					focus_x := row_i + i*dir[0]
					focus_y := col_i + i*dir[1]
					if is_out_of_range(focus_x, focus_y, len(row), len(table)) || table[focus_x][focus_y] != string(s) {
						break
					}
					if table[focus_x][focus_y] == "S" {
						found = true
					}
				}
				if found {
					count += 1
				}
			}
		}
	}
	fmt.Printf("answer: %d\n", count)
}

var DIRECTION_2 = [][]int{{-1, -1}, {-1, 1}}

// https://adventofcode.com/2024/day/4#part2
func partTwo(table [][]string) {
	count := 0
	for row_i, row := range table {
		for col_i, v := range row {
			// focus only A
			if v != "A" {
				continue
			}
			// find a valid direction
			count_word := 0
			for _, dir := range DIRECTION_2 {
				focus_x := row_i + dir[0]
				focus_y := col_i + dir[1]
				op_x := row_i - dir[0]
				op_y := col_i - dir[1]
				out_of_range := is_out_of_range(focus_x, focus_y, len(row), len(table)) || is_out_of_range(op_x, op_y, len(row), len(table))
				is_x_mas := false
				if !out_of_range {
					is_x_mas = strings.Contains("MS", table[focus_x][focus_y]) && strings.Contains("MS", table[op_x][op_y])
				}
				if out_of_range || table[focus_x][focus_y] == table[op_x][op_y] || !is_x_mas {
					break
				}
				count_word += 1
			}

			if count_word == 2 {
				count += 1
			}
		}
	}
	fmt.Printf("answer: %d\n", count)
}

func readFile(fileName string) string {
	b, err := os.ReadFile(fileName) // return in bytes
	if err != nil {
		fmt.Print(err)
		return ""
	}
	str := string(b)
	return str
}

func formatInput(longtxt string) [][]string {
	rows := strings.Split(longtxt, "\n")
	var table [][]string
	for _, v := range rows {
		col := strings.Split(v, "")
		table = append(table, col)
	}
	return table
}

func main() {

	table := formatInput(readFile("input.txt"))

	// start
	partTwo(table)

}
