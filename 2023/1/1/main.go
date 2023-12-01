package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"unicode"
)

func main() {
	texts := readInput()

	var first rune
	var second rune
	nums := make([]int, len(texts))

	for i, t := range texts {
		first = 0
		second = 0

		for _, c := range t {
			if unicode.IsDigit(c) {

				if first == 0 {
					first = c
					second = c
					continue
				}

				second = c

			}

		}
		nums[i] = int(first - '0')
		nums[i] = nums[i]*10 + int(second-'0')

	}

	sum := 0
	for _, n := range nums {
		sum += n
	}
	fmt.Println(sum)

}

// readInput reads the input.txt file and returns a slice of strings
// with each line of the file as an element of the slice.
func readInput() []string {

	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	open, err := os.Open(filepath.Join(currentDir, "2023/1/1/input.txt"))
	if err != nil {
		panic(err)
	}

	defer open.Close()
	var lines []string
	scanner := bufio.NewScanner(open)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines

}
