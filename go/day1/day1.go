package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Match struct {
	Position int
	Value    string
}

func parseMatches(line string) []Match {
	digits := [18]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	matches := []Match{}
	for _, digit := range digits {
		start := 0
		for {
			idx := strings.Index(line[start:], digit)
			if idx == -1 {
				// println("Not found", digit, "starting", start, "in", line[start:])
				break
			}
			// println("Found", digit, "at", start+idx)
			display_digit := digit
			if display_digit == "one" {
				display_digit = "1"
			} else if display_digit == "two" {
				display_digit = "2"
			} else if display_digit == "three" {
				display_digit = "3"
			} else if display_digit == "four" {
				display_digit = "4"
			} else if display_digit == "five" {
				display_digit = "5"
			} else if display_digit == "six" {
				display_digit = "6"
			} else if display_digit == "seven" {
				display_digit = "7"
			} else if display_digit == "eight" {
				display_digit = "8"
			} else if display_digit == "nine" {
				display_digit = "9"
			}

			matches = append(matches, Match{start + idx, display_digit})
			start = start + idx + 1
		}
	}
	sort.Slice(matches, func(i, j int) bool {
		return matches[i].Position < matches[j].Position
	})
	return matches
}

func main() {
	data, err := os.Open("day1/day1.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(data)
	calibration := 0
	for scanner.Scan() {
		line := scanner.Text()

		// part 2
		// scan for the first matching named number and replace it, repeat until there's no more matches
		matches := parseMatches(line)
		detected_number := matches[0].Value + matches[len(matches)-1].Value
		number, err := strconv.Atoi(detected_number)
		if err != nil {
			panic(err)
		}
		// debugging
		// if (line[0:1] + line[len(line)-1:]) != detected_number {
		// 	fmt.Printf("%+v %+v %+v\n", line, number, matches)
		// }
		calibration += number
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	println("Calibration", calibration)
}
