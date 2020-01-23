package main

import "testing"

func TestGetGuardId(t *testing.T) {
	line := "[1518-06-02 23:58] Guard #179 begins shift"
	id, err := get_guard(line)
	if err != nil {
		t.Error("getting guard id failed:", err)
	}

	if id != 179 {
		t.Error("got bad guard id:", id)
	}
}

func TestGetMinute(t *testing.T) {
	line := "[1518-07-26 23:57] Guard #1051 begins shift"
	min, err := get_minute(line)
	if err != nil {
		t.Error("getting minute failed:", err)
	}

	if min != 57 {
		t.Error("got bad minute:", min)
	}
}

func TestGetHour(t *testing.T) {
	line := "[1518-07-26 23:57] Guard #1051 begins shift"
	hour, err := get_hour(line)
	if err != nil {
		t.Error("getting minute failed:", err)
	}

	if hour != 23 {
		t.Error("got bad minute:", hour)
	}
}
