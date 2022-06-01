package util

import (
	"io/ioutil"
	"strings"
)

func ParseInputAsSlice(filename string) ([]string, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	contents := string(bytes)
	split := strings.Split(contents, "\n")
	return split, nil
}
