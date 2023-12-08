package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func validSerialNumber(grid [][]string, num_text string, row_id int, col_id int) bool {
	// we have the position of the last character of the number
	// we need to check each character in the number, every direction but the string itself
	// if it's anything other than a . then it's valid!
	num_text_len := len(num_text)
	col_id_start := col_id - num_text_len
	for i := col_id_start; i <= row_id; i++ {

	}
	return true
}

func main() {
	data, err := os.ReadFile("day3/day3.txt")
	if err != nil {
		panic(err)
	}
	grid := [][]string{}
	sum := 0
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		row := make([]string, 0)
		for _, char := range line {
			row = append(row, string(char))
		}
		grid = append(grid, row)
	}

	for row_id, row := range grid {
		current_num := ""
		for col_id, char := range row {
			// if char is number put it in current_num
			_, err := strconv.Atoi(char)
			if err == nil {
				current_num += char
			} else if current_num != "" {
				// got a full number!
				if validSerialNumber(grid, current_num, row_id, col_id-1) {
					fmt.Printf("validSerialNumber %d,%d %s\n", row_id, col_id-1, current_num)
					serial, err := strconv.Atoi(current_num)
					if err == nil {
						sum += serial
					}
				}
				current_num = ""
			} else {
				current_num = ""
			}
			// fmt.Printf("%d,%d %s\n", row_id, col_id, char)
		}

		// end of row!
		if current_num != "" {
			// got a full number!
			if validSerialNumber(grid, current_num, row_id, len(row)-1) {
				serial, err := strconv.Atoi(current_num)
				if err == nil {
					sum += serial
				}
			}
		}
	}

	// fmt.Printf("%+v\n", grid)
	fmt.Printf("Grid is %d x %d\n", len(grid), len(grid[0]))
	fmt.Printf("Sum is %d\n", sum)
}
