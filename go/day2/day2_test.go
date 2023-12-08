package main

import (
	"reflect"
	"testing"
)

func TestParseLine(t *testing.T) {
	line := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	game := parseLine(line)
	want := Game{
		id: 1,
		rounds: []Bag{
			{red: 4, green: 0, blue: 3},
			{red: 1, green: 2, blue: 6},
			{red: 0, green: 2, blue: 0},
		},
	}
	if !reflect.DeepEqual(game, want) {
		t.Fatalf(`parseLine("") = %+v, want %+v`, game, want)
	}
}

func TestParseLine2(t *testing.T) {
	line := "Game 24: 5 blue, 15 green, 13 red; 20 green, 13 blue, 6 red; 5 blue, 11 red, 16 green; 6 red, 5 blue, 13 green; 12 blue, 13 green, 3 red"
	game := parseLine(line)
	want := Game{
		id: 24,
		rounds: []Bag{
			{red: 13, green: 15, blue: 5},
			{red: 6, green: 20, blue: 13},
			{red: 11, green: 16, blue: 5},
			{red: 6, green: 13, blue: 5},
			{red: 3, green: 13, blue: 12},
		},
	}
	if !reflect.DeepEqual(game, want) {
		t.Fatalf(`parseLine("") = %+v, want %+v`, game, want)
	}
}

func TestValidGame(t *testing.T) {
	bag := Bag{
		red:   13,
		green: 20,
		blue:  15,
	}
	game := Game{
		id: 24,
		rounds: []Bag{
			{red: 13, green: 15, blue: 5},
			{red: 6, green: 20, blue: 13},
			{red: 11, green: 16, blue: 5},
			{red: 6, green: 13, blue: 5},
			{red: 3, green: 13, blue: 12},
		},
	}
	if !validGame(bag, game) {
		t.Fatalf(`validGame(bag, game) = false when it should be true`)
	}
}

func TestInvalidGame(t *testing.T) {
	bag := Bag{
		red:   10,
		green: 15,
		blue:  20,
	}
	game := Game{
		id: 24,
		rounds: []Bag{
			{red: 13, green: 15, blue: 5},
			{red: 6, green: 20, blue: 13},
			{red: 11, green: 16, blue: 5},
			{red: 6, green: 13, blue: 5},
			{red: 3, green: 13, blue: 12},
		},
	}
	if validGame(bag, game) {
		t.Fatalf(`validGame(bag, game) = true when it should be false`)
	}
}
