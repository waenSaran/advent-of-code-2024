package main

import (
	"fmt"
	"os"
	"strings"
)

var STRING = "XMAS"
var LEN_STRING = len(STRING)
var DIRECTION = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

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
					out_of_range := focus_x < 0 || focus_y < 0 || focus_x >= len(row) || focus_y >= len(table)
					if out_of_range || table[focus_x][focus_y] != string(s) {
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
	partOne(table)

}
