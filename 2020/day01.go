package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = "inputs/day01.txt"

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(contents, "\n")

	// part A
	for _, s := range split {
		i, _ := strconv.Atoi(s)
		fmt.Println(i)
	}
}
