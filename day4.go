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

	guard_map := analyse_log()

	sleepiest_guard := -1
	mins_asleep := -1
	worst_minute := -1

	for id, guard_info := range guard_map {
		total_mins_asleep := 0
		var most_days int
		var worst_min int
		for min, days := range guard_info {
			if days > most_days {
				most_days = days
				worst_min = min
			}
			total_mins_asleep += days
		}

		if total_mins_asleep > mins_asleep {
			sleepiest_guard = id
			mins_asleep = total_mins_asleep
			worst_minute = worst_min
		}
	}

	return fmt.Sprintf("%d * %d = %d", sleepiest_guard, worst_minute, sleepiest_guard*worst_minute), nil
}

func day4_2() (result string, err error) {
	guard_map := analyse_log()

	var guard int
	var minute int
	var peak_count int

	for id, guard_info := range guard_map {
		for min, count := range guard_info {
			if count > peak_count {
				guard = id
				minute = min
				peak_count = count
			}
		}
	}

	return fmt.Sprintf("%d * %d = %d", guard, minute, guard*minute), nil
}

func analyse_log() (guard_map GuardMap) {

	data, err := read_lines("input/day4.txt")
	if err != nil {
		return
	}

	sort.Strings(data)
	guard_map = make(GuardMap)
	mode := 0
	current_guard := -1
	asleep_time := -1
	for _, line := range data {

		var hr int
		hr, err = get_hour(line)
		if err != nil {
			return
		}

		switch mode {
		case 0: // initial state
			if strings.Contains(line, "Guard") {
				current_guard, err = get_guard(line)
				if err != nil {
					return
				}
			} else if strings.Contains(line, "falls asleep") {
				if hr != 0 {
					continue
				}
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
				if hr != 0 {
					continue
				}
				var wake_time int
				wake_time, err = get_minute(line)
				if err != nil {
					return
				}

				i := 0
				for ; i < wake_time-asleep_time; i++ {
					_, ok := guard_map[current_guard][i+asleep_time]
					if ok {
						guard_map[current_guard][i+asleep_time] += 1
					} else {
						if guard_map[current_guard] == nil {
							guard_map[current_guard] = make(Schedule)
						}
						guard_map[current_guard][i+asleep_time] = 1
					}
				}
				asleep_time = -1
				mode = 0
			}
		}
	}

	return
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
