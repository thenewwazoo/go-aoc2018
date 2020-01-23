package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Schedule map[int]int
type GuardMap map[int]Schedule

func day4_1() (result string, err error) {
	data, err := read_lines("input/day4.txt")
	if err != nil {
		return
	}

	sort.Strings(data)

	guard_map := make(GuardMap)
	mode := 0
	current_guard := -1
	asleep_time := -1
	for _, line := range data {
		hr, err = get_hour(line)
		if err != nil {
			return
		}
		if hr != 0 {
			continue
		}
		switch mode {
		case 0: // initial state
			if strings.Contains(line, "Guard") {
				current_guard, err = get_guard(line)
				if err != nil {
					return
				}
				mode = 1
			} else if strings.Contains(line, "falls asleep") {
				asleep_time, err = get_minute(line)
				if err != nil {
					return
				}
				mode = 1
			} else {
				panic(fmt.Sprintf("got unexpected line: %s", line))
			}
		case 1: // asleep
			if strings.Contains(line, "wakes up") {
				var wake_time int
				wake_time, err = get_minute(line)
				if err != nil {
					return
				}

				for i := asleep_time; i < wake_time; i++ {
					_, ok := guard_map[current_guard][i]
					if ok {
						guard_map[current_guard][i] += 1
					} else {
						if guard_map[current_guard] == nil {
							guard_map[current_guard] = make(Schedule)
						}
						guard_map[current_guard][i] = 1
					}
				}
				asleep_time = -1
				mode = 0
			}
		}
	}

	sleepiest_guard := -1
	mins_asleep := -1
	worst_minute := -1
	for id, guard_info := range guard_map {
		if len(guard_info) > mins_asleep {
			sleepiest_guard = id
			mins_asleep = len(guard_info)

			worst_minute = -1
			highest_found := 0
			for minute, slept := range guard_info {
				if slept > highest_found {
					highest_found = slept
					worst_minute = minute
				}
			}
		}
	}

	return fmt.Sprintf("%d * %d = %d", sleepiest_guard, worst_minute, sleepiest_guard*worst_minute), nil
}

func get_guard(line string) (id int, err error) {
	engine := regexp.MustCompile(`#(\d+)`)
	parts := engine.FindStringSubmatch(line)[1:]
	_id, err := strconv.ParseInt(parts[0], 10, 64)
	id = int(_id)
	return
}

func get_minute(line string) (min int, err error) {
	engine := regexp.MustCompile(`:(\d\d)]`)
	parts := engine.FindStringSubmatch(line)[1:]
	_min, err := strconv.ParseInt(parts[0], 10, 64)
	min = int(_min)
	return
}

func get_hour(line string) (hour int, err error) {
	engine := regexp.MustCompile(` (\d\d):`)
	parts := engine.FindStringSubmatch(line)[1:]
	_hour, err := strconv.ParseInt(parts[0], 10, 64)
	hour = int(_hour)
	return
}
