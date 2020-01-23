package main

import (
	"fmt"
	"os"
)

func runit(day int, f func() (string, error)) {
	r, err := f()
	if err != nil {
		fmt.Printf("day %d err: %s\n", day, err)
	} else {
		fmt.Printf("day %d result: %s\n", day, r)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("need day")
		return
	}
	day := os.Args[1]

	switch day {
	case "1":
		runit(1, day1_1)
		runit(1, day1_2)
	case "2":
		runit(2, day2_1)
		runit(2, day2_2)
	case "3":
		runit(3, day3_1)
		runit(3, day3_2)
	case "4":
		runit(4, day4_1)
	default:
		fmt.Println("bleh")
	}

}
