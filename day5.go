package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
	"unicode"
)

func day5_1() (result string, err error) {

	fh, err := os.Open("input/day5.txt")
	if err != nil {
		return
	}

	r := bufio.NewReader(fh)

	l := list.New()
	c_map := make(map[rune]bool)
	for {
		var c rune
		var n int
		c, n, err = r.ReadRune()
		C := unicode.ToUpper(c)
		if err != nil || n == 0 {
			break
		}

		l.PushBack(c)
		_, ok := c_map[C]
		if !ok {
			c_map[C] = true
		}
	}

	l = react_chain(l)

	return fmt.Sprint(l.Len()), nil
}

func is_opp_pol(l rune, r rune) bool {
	if unicode.IsUpper(l) {
		if unicode.IsLower(r) {
			if unicode.ToUpper(l) == unicode.ToUpper(r) {
				return true
			}
		}
	} else {
		if unicode.IsUpper(r) {
			if unicode.ToUpper(l) == unicode.ToUpper(r) {
				return true
			}
		}
	}
	return false
}

func react_chain(l *list.List) *list.List {
	out := list.New()

	out.PushBack(l.Front().Value.(rune))
	for e := l.Front().Next(); e != nil; e = e.Next() {
		y := e.Value.(rune)

		if out.Front() == nil {
			//fmt.Println("out is empty: pushing ", string(y))
			out.PushBack(y)
			continue
		}
		x := out.Back().Value.(rune)

		//fmt.Println("checking", string(x), string(y))

		if is_opp_pol(x, y) {
			//fmt.Println("match, removing ", string(x))
			out.Remove(out.Back())
		} else {
			//fmt.Println("no match, pushing ", string(y))
			out.PushBack(y)
		}

		//fmt.Print("current out: ")
		//for o := out.Front(); o != nil; o = o.Next() {
		//fmt.Print(string(o.Value.(rune)))
		//}
		//fmt.Println()

	}
	return out
}

func remove_all(x rune, l *list.List) {
	for e := l.Front(); e != nil; {
		v := e.Value.(rune)
		if unicode.ToUpper(v) == unicode.ToUpper(x) {
			if e.Next() != nil {
				e = e.Next()
				l.Remove(e.Prev())
			} else {
				l.Remove(e)
				break
			}
		} else {
			e = e.Next()
		}
	}
}

func day5_2() (result string, err error) {

	fh, err := os.Open("input/day5.txt")
	if err != nil {
		return
	}

	r := bufio.NewReader(fh)

	l := list.New()
	c_map := make(map[rune]bool)
	for {
		var c rune
		var n int
		c, n, err = r.ReadRune()
		C := unicode.ToUpper(c)
		if err != nil || n == 0 {
			break
		}

		l.PushBack(c)
		_, ok := c_map[C]
		if !ok {
			c_map[C] = true
		}
	}

	best_len := math.MaxInt32
	for c, _ := range c_map {
		m := list.New()

		m.PushBackList(l)

		remove_all(c, m)
		m = react_chain(m)

		if m.Len() < best_len {
			best_len = m.Len()
		}
	}

	return fmt.Sprint(best_len), nil
}
