package main

import (
	"reflect"
	"testing"
)

func TestParseClaim(t *testing.T) {
	var line = "#123 @ 3,2: 5x4"
	var expected = Claim{
		id:       123,
		x_offset: 3,
		y_offset: 2,
		width:    5,
		height:   4,
	}
	r, err := parse_claim(line)
	if err != nil {
		t.Errorf("bad parse got err: %s", err)
		return
	}

	if !reflect.DeepEqual(r, expected) {
		t.Errorf("got %#v expected %#v", r, expected)
	}
}

func TestMapOverlap(t *testing.T) {

	// #1 @ 1,3: 4x4
	var first = Claim{
		id:       1,
		x_offset: 1,
		y_offset: 3,
		width:    4,
		height:   4,
	}

	// #2 @ 3,1: 4x4
	var second = Claim{
		id:       2,
		x_offset: 3,
		y_offset: 1,
		width:    4,
		height:   4,
	}

	// #3 @ 5,5: 2x2
	var third = Claim{
		id:       3,
		x_offset: 5,
		y_offset: 5,
		width:    2,
		height:   2,
	}

	var first_third = make(map[Coord]bool)

	var first_second = map[Coord]bool{
		Coord{x: 3, y: 3}: true,
		Coord{x: 3, y: 4}: true,
		Coord{x: 4, y: 3}: true,
		Coord{x: 4, y: 4}: true,
	}

	r := map_overlap(first, third)
	if !reflect.DeepEqual(r, first_third) {
		t.Error("bad overlap 1 v 3:", r)
		return
	}

	r = map_overlap(first, second)
	if !reflect.DeepEqual(r, first_second) {
		t.Error("bad overlap 1 v 2:", r)
		return
	}
}

func TestIsInClaim(t *testing.T) {

	// #1 @ 1,3: 4x4
	var c = Claim{
		id:       1,
		x_offset: 1,
		y_offset: 3,
		width:    4,
		height:   4,
	}
	// o....... -> x
	// ........
	// ........
	// .1111...
	// .1111...
	// .1111...
	// .1111...
	// ........
	// |
	// v
	// y

	var in = Coord{
		x: 1,
		y: 4,
	}
	var out = Coord{
		x: 5,
		y: 7,
	}

	if !is_in_claim(in, c) {
		t.Error("failed to find", in, "in", c)
		return
	}

	if is_in_claim(out, c) {
		t.Error("erroneously found", out, "in", c)
	}
}

func TestMergeMaps(t *testing.T) {

	var first = map[Coord]bool{
		Coord{x: 4, y: 3}: true,
		Coord{x: 4, y: 4}: true,
	}

	var second = map[Coord]bool{
		Coord{x: 3, y: 3}: true,
		Coord{x: 3, y: 4}: true,
	}

	var result = map[Coord]bool{
		Coord{x: 3, y: 3}: true,
		Coord{x: 3, y: 4}: true,
		Coord{x: 4, y: 3}: true,
		Coord{x: 4, y: 4}: true,
	}

	output := merge_maps(first, second)
	if !reflect.DeepEqual(result, output) {
		t.Error("merged map is ", output, "expected", result)
		return
	}

}
