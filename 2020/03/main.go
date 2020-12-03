package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)
var right, down int
func init() {
	flag.IntVar(&right, "r", 1, "Right Increment Value")
	flag.IntVar(&down, "d", 2, "Down Increment Value")
}

func growLine(line string, i int) string {
	length := len(line)
	for length <= i {
		line = line + line
		length = len(line)
	}
	return line
}

func main() {
	flag.Parse()
	f, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(f)
	i := 0
	l := 0
	n := 0
	for scanner.Scan() {
		var line string
		line = scanner.Text()
		line = growLine(line, i)
		if line[i] == '#' {
			fmt.Printf("line %v, position %v: %v\n", l, i, line)
			n++
		}
		i += right
		l += down
		if down > 1 {
			for j := 0; j < (down - 1); j++ {
				scanner.Scan()
			}
		}
	}
	fmt.Println(n)
}
