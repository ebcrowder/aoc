package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = "inputs/day02.txt"

func main() {
	bytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(contents, "\n")
	split = split[:len(split)-1] // trim blank element at end of list

	// part A
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

	fmt.Println("Part A:", product)

	// part B
	horizontal = 0
	depth = 0
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

	product = horizontal * depth

	fmt.Println("Part B:", product)
}
