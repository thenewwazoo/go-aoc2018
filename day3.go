package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type Claim struct {
	id       uint64
	x_offset uint64
	y_offset uint64
	width    uint64
	height   uint64
}

type Coord struct {
	x uint64
	y uint64
}

func is_in_claim(c Coord, m Claim) bool {
	if c.x >= m.x_offset && c.x < m.x_offset+m.width && c.y >= m.y_offset && c.y < m.y_offset+m.height {
		return true
	} else {
		return false
	}
}

func map_claim(l Claim) (result map[Coord]bool) {
	result = make(map[Coord]bool)
	for x := l.x_offset; x < l.x_offset+l.width; x++ {
		for y := l.y_offset; y < l.y_offset+l.height; y++ {
			loc := Coord{x: x, y: y}
			result[loc] = true
		}
	}
	return
}

func map_overlap(l Claim, r Claim) (result map[Coord]bool) {
	result = make(map[Coord]bool)

	for x := l.x_offset; x < l.x_offset+l.width; x++ {
		for y := l.y_offset; y < l.y_offset+l.height; y++ {
			loc := Coord{x: x, y: y}
			if is_in_claim(loc, r) {
				result[loc] = true
			}
		}
	}
	return
}

func merge_maps(l map[Coord]bool, r map[Coord]bool) map[Coord]bool {
	for k, v := range r {
		l[k] = v
	}
	return l
}

func parse_claim(raw string) (result Claim, err error) {
	// like "#123 @ 3,2: 5x4"
	engine := regexp.MustCompile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`)
	parts := engine.FindStringSubmatch(raw)[1:]

	result.id, err = strconv.ParseUint(string(parts[0]), 10, 64)
	if err != nil {
		fmt.Printf("bleh %#v", parts)
		return
	}

	result.x_offset, err = strconv.ParseUint(string(parts[1]), 10, 64)
	if err != nil {
		return
	}

	result.y_offset, err = strconv.ParseUint(string(parts[2]), 10, 64)
	if err != nil {
		return
	}

	result.width, err = strconv.ParseUint(string(parts[3]), 10, 64)
	if err != nil {
		return
	}

	result.height, err = strconv.ParseUint(string(parts[4]), 10, 64)
	if err != nil {
		return
	}

	return
}

func day3_1() (result string, err error) {

	data, err := read_lines("input/day3.txt")
	if err != nil {
		return
	}

	var claims []Claim
	for _, datum := range data {
		var c Claim
		c, err = parse_claim(datum)
		if err != nil {
			return
		}

		claims = append(claims, c)
	}

	total_map := make(map[Coord]int)
	for _, claim := range claims {
		c_map := map_claim(claim)
		for k, _ := range c_map {
			c, ok := total_map[k]
			if ok {
				total_map[k] = c + 1
			} else {
				total_map[k] = 1
			}
		}
	}

	for k, v := range total_map {
		if v <= 1 {
			delete(total_map, k)
		}
	}
	return fmt.Sprint(len(total_map)), nil
}

func day3_2() (result string, err error) {

	data, err := read_lines("input/day3.txt")
	if err != nil {
		return
	}

	var claims []Claim
	for _, datum := range data {
		var c Claim
		c, err = parse_claim(datum)
		if err != nil {
			return
		}

		claims = append(claims, c)
	}

	total_map := make(map[Coord]int)
	for _, claim := range claims {
		c_map := map_claim(claim)
		for k, _ := range c_map {
			c, ok := total_map[k]
			if ok {
				total_map[k] = c + 1
			} else {
				total_map[k] = 1
			}
		}
	}

	for _, claim := range claims {
		c_map := map_claim(claim)
		has_overlap := false
		for k, _ := range c_map {
			if total_map[k] > 1 {
				has_overlap = true
				break
			}
		}

		if !has_overlap {
			return fmt.Sprint(claim.id), nil
		}
	}
	err = errors.New("not found")
	return
}
