package main

import (
	"bufio"
	"os"
)

func calculateSmallestPossibleBag(game Game) Bag {
	bag := Bag{0, 0, 0}
	for _, round := range game.rounds {
		if round.blue > bag.blue {
			bag.blue = round.blue
		}
		if round.green > bag.green {
			bag.green = round.green
		}
		if round.red > bag.red {
			bag.red = round.red
		}
	}

	return bag
}

func day2b() {
	data, err := os.Open("day2/day2.txt")
	if err != nil {
		panic(err)
	}

	sum := 0
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		game := parseLine(line)
		bag := calculateSmallestPossibleBag(game)
		power := bag.blue * bag.green * bag.red
		sum += power
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	println("sum", sum)
}
