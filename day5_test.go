package main

import (
	"container/list"
	"fmt"
	"testing"
)

func TestIsOppPol(t *testing.T) {

	l := list.New()
	l.PushBack(rune('a'))
	l.PushBack(rune('A'))
	if !is_opp_pol(l.Front().Value.(rune), l.Back().Value.(rune)) {
		t.Error("did not find aA to be opp pol")
		return
	}

	l = list.New()
	l.PushBack(rune('a'))
	l.PushBack(rune('a'))
	if is_opp_pol(l.Front().Value.(rune), l.Back().Value.(rune)) {
		t.Error("found aa to be opp pol")
		return
	}

	l = list.New()
	l.PushBack(rune('a'))
	l.PushBack(rune('b'))
	if is_opp_pol(l.Front().Value.(rune), l.Back().Value.(rune)) {
		t.Error("found ab to be opp pol")
		return
	}

	l = list.New()
	l.PushBack(rune('a'))
	l.PushBack(rune('B'))
	if is_opp_pol(l.Front().Value.(rune), l.Back().Value.(rune)) {
		t.Error("found aB to be opp pol")
		return
	}

}

func TestReactChain(t *testing.T) {
	l := list.New()
	l.PushBack(rune('a'))
	l.PushBack(rune('b'))
	l.PushBack(rune('B'))
	l.PushBack(rune('c'))

	l = react_chain(l)

	if l.Front().Value.(rune) != 'a' || l.Front().Next().Value.(rune) != 'c' {
		t.Error("did not collapse abBc to ac")
		return
	}

	l = list.New()
	l.PushBack('a')
	l.PushBack('b')
	l.PushBack('c')
	l.PushBack('C')
	l.PushBack('B')
	l.PushBack('d')

	l = react_chain(l)

	if l.Front().Value.(rune) != 'a' || l.Front().Next().Value.(rune) != 'd' {
		t.Error("did not collapse abcCBd to ad")
		return
	}

}

func TestRemoveAll(t *testing.T) {

	l := list.New()
	l.PushBack('a')
	l.PushBack('b')
	l.PushBack('c')
	l.PushBack('C')
	l.PushBack('B')
	l.PushBack('d')

	remove_all('c', l)
	remove_all('b', l)

	if l.Front().Value.(rune) != 'a' || l.Front().Next().Value.(rune) != 'd' {
		t.Error("did not remove c and d to leave ad")
		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Print(string(e.Value.(rune)))
		}
		fmt.Println()
		return
	}

	l = list.New()
	l.PushBack('a')
	l.PushBack('b')
	l.PushBack('B')
	l.PushBack('A')

	remove_all('a', l)

	if l.Front().Value.(rune) != 'b' || l.Front().Next().Value.(rune) != 'B' {
		t.Error("did not remove as to leave bB")
		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Print(string(e.Value.(rune)))
		}
		fmt.Println()
		return
	}

}
