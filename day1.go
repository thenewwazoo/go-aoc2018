package main

import (
	"fmt"
	"strconv"
)

func day1_1() (result string, err error) {
	var data []string
	data, err = read_lines("input/day1.txt")
	if err != nil {
		return
	}
	var counter = 0

	for _, datum := range data {
		var val int
		val, err = strconv.Atoi(datum)
		if err != nil {
			return
		}
		counter += val
	}
	return fmt.Sprintf("%d", counter), nil
}

func day1_2() (result string, err error) {
	data, err := read_lines("input/day1.txt")
	if err != nil {
		return
	}

	freqmap := make(map[int]bool)
	var counter = 0
	found := false
	for found == false {
		for _, datum := range data {
			var val int
			val, err = strconv.Atoi(datum)
			if err != nil {
				return
			}
			counter += val
			_, dupe := freqmap[counter]
			if dupe {
				found = true
				break
			}
			freqmap[counter] = true
		}
	}

	return fmt.Sprintf("%d", counter), nil
}
