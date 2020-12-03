package main

import (
	"bufio"
	"fmt"
	"os"
)

func growLine(line string, i int) string {
	length := len(line)
	for length <= i {
		line = line + line
		length = len(line)
	}
	return line
}

func main() {
	f, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(f)
	i := 0
	l := 0
	n := 0
	for scanner.Scan() {
		i += 3
		l++
		line := scanner.Text()
		line = growLine(line, i)
		if line[i] == '#' {
			fmt.Printf("line %v, position %v: %v\n", l, i, line)
			n++
		}
	}
	fmt.Println(n)
}
