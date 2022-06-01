package main

import (
	"strconv"
	"strings"
	"testing"

	"github.com/ebcrowder/aoc/v2/lib/util"
)

func TestDay02PartA(t *testing.T) {
	input := "inputs/day02.txt"
	split, _ := util.ParseInputAsSlice(input)
	split = split[:len(split)-1] // trim blank element at end of list

	horizontal := 0
	depth := 0

	for _, direction := range split {
		s := strings.Split(direction, " ")
		i, _ := strconv.Atoi(s[1])
		switch s[0] {
		case "up":
			depth -= i
		case "down":
			depth += i
		case "forward":
			horizontal += i
		}
	}

	product := horizontal * depth

	want := 2039912
	if product != want {
		t.Fatalf("got %q, want %#q", product, want)
	}
}

func TestDay02PartB(t *testing.T) {
	input := "inputs/day02.txt"
	split, _ := util.ParseInputAsSlice(input)
	split = split[:len(split)-1] // trim blank element at end of list

	horizontal := 0
	depth := 0
	aim := 0

	for _, direction := range split {
		s := strings.Split(direction, " ")
		i, _ := strconv.Atoi(s[1])
		switch s[0] {
		case "up":
			aim -= i
		case "down":
			aim += i
		case "forward":
			horizontal += i
			depth += aim * i
		}
	}

	product := horizontal * depth

	want := 1942068080
	if product != want {
		t.Fatalf("got %q, want %#q", product, want)
	}
}
