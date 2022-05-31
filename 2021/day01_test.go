package main

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

var input = "inputs/day01.txt"

func parseInput() ([]string, error) {
	bytes, err := ioutil.ReadFile(input)
	if err != nil {
		return nil, err
	}
	contents := string(bytes)
	split := strings.Split(contents, "\n")
	return split, nil
}

func partA(split []string) int {
	increments := 0
	last := -1
	for _, s := range split {
		i, _ := strconv.Atoi(s)
		if i > last && last != -1 {
			increments++
		}
		last = i

	}
	return increments
}

func partB(split []string) int {
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
	return increments
}

func TestPartA(t *testing.T) {
	split, _ := parseInput()

	want := 1154
	partA := partA(split)
	if want != partA {
		t.Fatalf("got %q, want %#q", partA, want)
	}
}

func TestPartB(t *testing.T) {
	split, _ := parseInput()

	want := 1127
	partB := partB(split)
	if want != partB {
		t.Fatalf("got %q, want %#q", partB, want)
	}
}
