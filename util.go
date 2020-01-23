package main

import (
	"bufio"
	"os"
	"strings"
)

func read_lines(filename string) (lines []string, err error) {
	fh, err := os.Open(filename)
	if err != nil {
		return
	}
	defer fh.Close()

	scanner := bufio.NewScanner(fh)

	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}

	return
}
