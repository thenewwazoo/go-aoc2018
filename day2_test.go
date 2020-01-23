package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestDist(t *testing.T) {
	var test_data = []struct {
		s string
		t string
		c []byte
		d int
	}{
		{
			"abcde", "axcye", []byte{'a', 'c', 'e'}, 2,
		},
		{
			"fghij", "fguij", []byte{'f', 'g', 'i', 'j'}, 1,
		},
	}

	for _, datum := range test_data {
		commons, count, err := find_commons(datum.s, datum.t)
		if err != nil {
			t.Errorf("Err! %s", err)
			return
		}
		if bytes.Equal(commons, datum.c) && count != datum.d {
			t.Error("No match, got ", count, commons, "expected ", datum.d, datum.c)
			return
		}
	}
}

func TestCountem(t *testing.T) {
	var test_list = []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}
	twos, threes := count_twosthrees(test_list)
	if twos != 4 && threes != 3 {
		t.Error("Bad counts: doubles: {} triples: {}", twos, threes)
	}
}

func TestStrHist(t *testing.T) {
	var test_data = []struct {
		s string
		m map[byte]int
	}{
		{
			"abcdef",
			map[byte]int{
				'a': 1,
				'b': 1,
				'c': 1,
				'd': 1,
				'e': 1,
				'f': 1,
			},
		},
		{
			"aabcde",
			map[byte]int{
				'a': 2,
				'b': 1,
				'c': 1,
				'd': 1,
				'e': 1,
			},
		},
	}
	for _, datum := range test_data {
		hist := str_hist(datum.s)
		if !reflect.DeepEqual(hist, datum.m) {
			t.Error("bleh")
			break
		}
	}
}
