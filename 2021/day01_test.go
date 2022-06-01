package main

import (
	"strconv"
	"testing"

	"github.com/ebcrowder/aoc/v2/lib/util"
)

func TestDay01PartA(t *testing.T) {
	input := "inputs/day01.txt"
	split, _ := util.ParseInputAsSlice(input)

	increments := 0
	last := -1
	for _, s := range split {
		i, _ := strconv.Atoi(s)
		if i > last && last != -1 {
			increments++
		}
		last = i

	}
	want := 1154
	if increments != want {
		t.Fatalf("got %q, want %#q", increments, want)
	}
}

func TestDay01PartB(t *testing.T) {
	input := "inputs/day01.txt"
	split, _ := util.ParseInputAsSlice(input)

	increments := 0
	lastNums := []int{0, 0, 0}
	lastSum := 0

	for idx, s := range split {
		i, _ := strconv.Atoi(s)
		sum := i + lastSum - lastNums[0]
		lastNums = append(lastNums[1:], i)
		if idx < 3 {
			lastSum = sum
			continue
		}
		if sum > lastSum {
			increments++
		}
		lastSum = sum
	}

	want := 1127
	if increments != want {
		t.Fatalf("got %q, want %#q", increments, want)
	}
}
