package main

import (
	"bufio"
	"os"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var r = regexp.MustCompile(`(\d+)-(\d+)\s([a-z]):\s([a-z]*)`)

func isValidInput(line string) bool {
	groups := r.FindStringSubmatch(line)
	lowest, _ := strconv.Atoi(groups[1])
	highest, _ := strconv.Atoi(groups[2])
	target := groups[3]
	input := groups[4]
	count := strings.Count(input, target)
	if count >= lowest && count <= highest {
		return true
	}
	return false
}

func main() {
	f, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(f)
	n := 0
	for scanner.Scan() {
		line := scanner.Text()
		if isValidInput(line) {
			n++
		}
	}
	fmt.Println(n)
}
