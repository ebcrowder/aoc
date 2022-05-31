package template

import (
	"flag"
	"fmt"
	"io/ioutil"
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

	// TODO
	fmt.Println(split)

	// part A

	// part B
}
