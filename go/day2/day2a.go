package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Bag struct {
	red   int
	green int
	blue  int
}

type Game struct {
	id     int
	rounds []Bag
}

func parseLine(line string) Game {
	gameRegex := regexp.MustCompile(`^Game (\d+): (.+)`)
	cubeRegex := regexp.MustCompile(`(\d+) (.+)`)
	matches := gameRegex.FindStringSubmatch(line)
	if matches == nil {
		panic("Didn't match line " + line)
	}
	game_id, err := strconv.Atoi(matches[1])
	if err != nil {
		panic(err)
	}
	var rounds []Bag
	round_strings := strings.Split(matches[2], "; ")
	for _, round_string := range round_strings {
		cubeStrings := strings.Split(round_string, ", ")
		round := Bag{0, 0, 0}
		for _, cube := range cubeStrings {
			matches := cubeRegex.FindStringSubmatch(cube)
			if matches == nil {
				panic("didn't match cube " + cube)
			}
			cube_count, err := strconv.Atoi(matches[1])
			if err != nil {
				panic(err)
			}
			color := matches[2]
			if color == "red" {
				round.red = cube_count
			} else if color == "green" {
				round.green = cube_count
			} else if color == "blue" {
				round.blue = cube_count
			} else {
				panic("Unknown color " + color)
			}
		}
		rounds = append(rounds, round)
	}
	return Game{
		game_id,
		rounds,
	}
}

func validGame(bag Bag, game Game) bool {
	for _, game := range game.rounds {
		if game.red > bag.red || game.green > bag.green || game.blue > bag.blue {
			return false
		}
	}
	return true
}

func day2a() {
	data, err := os.Open("day2/day2.txt")
	if err != nil {
		panic(err)
	}
	bag := Bag{red: 12, green: 13, blue: 14}
	sum := 0
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		game := parseLine(line)
		if validGame(bag, game) {
			sum += game.id
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	println("sum", sum)
}
