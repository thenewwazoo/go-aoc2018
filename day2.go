package main

import "fmt"
import "errors"

func find_commons(l string, r string) (result []byte, count int, err error) {
	if len(l) != len(r) {
		err = errors.New(fmt.Sprintf("lengths do not match: %s(%d) != %s(%d)", l, len(l), r, len(r)))
		return
	}
	for i, c := range []byte(l) {
		d := r[i]
		if c != d {
			count += 1
		} else {
			result = append(result, c)
		}
	}
	return
}

func count_twosthrees(data []string) (int, int) {
	twos_count := 0
	threes_count := 0

	for _, s := range data {
		has_two, has_three := false, false
		for _, n := range str_hist(s) {
			if n == 2 {
				has_two = true
			} else if n == 3 {
				has_three = true
			}
		}
		if has_two {
			twos_count += 1
		}
		if has_three {
			threes_count += 1
		}
	}
	return twos_count, threes_count
}

func str_hist(s string) (out map[byte]int) {
	out = make(map[byte]int)
	for _, c := range []byte(s) {
		_, ok := out[c]
		if !ok {
			out[c] = 1
		} else {
			out[c] += 1
		}
	}
	return
}

func day2_1() (result string, err error) {
	data, err := read_lines("input/day2.txt")
	if err != nil {
		return
	}

	twos, threes := count_twosthrees(data)
	return fmt.Sprintf("%d * %d = %d", twos, threes, twos*threes), nil
}

func day2_2() (result string, err error) {
	data, err := read_lines("input/day2.txt")
	if err != nil {
		return
	}

	for i, s := range data {
		for _, t := range data[i+1:] {
			var commons []byte
			var count int
			commons, count, err = find_commons(s, t)
			if err != nil {
				return
			}

			if count == 1 {
				result = fmt.Sprintf("%s", string(commons))
				return
			}
		}
	}
	return
}
