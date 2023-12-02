package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"unicode"
)

func main() {
	texts := readInput("2023/1/input.txt")
	part1(texts)
	part2(texts)

}

func part1(input []string) {
	var first rune
	var second rune
	nums := make([]int, len(input))

	sum := 0

	for i, t := range input {
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
		nums[i] = convertRuneToInt(first, second)
		sum += nums[i]

	}
	fmt.Println(sum)
}

func part2(input []string) {
	var first rune
	var second rune
	nums := make([]int, len(input))

	numbers := map[string]rune{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}
	sum := 0

	for i, t := range input {
		matches := getMatches(t)
		index := 0
		first = 0
		second = 0
		for _, match := range matches {
			var c rune
			for match[0] > index {
				c = rune(t[index])
				if unicode.IsDigit(c) {

					if first == 0 {
						first = c
						second = c
						index++
						continue
					}

					second = c

				}
				index++
			}

			if first == 0 {
				first = numbers[t[match[0]:match[1]]]
				second = numbers[t[match[0]:match[1]]]
				index = match[1]
				continue
			}

			second = numbers[t[match[0]:match[1]]]
			index = match[1]
		}

		if index <= len(t)-1 {
			for _, c := range t[index:] {
				if unicode.IsDigit(c) {

					if first == 0 {
						first = c
						second = c
						continue
					}

					second = c

				}

			}
		}

		nums[i] = convertRuneToInt(first, second)
		sum += nums[i]

	}

	fmt.Println(sum)

}

func convertRuneToInt(first, second rune) int {
	num := int(first - '0')
	num = num*10 + int(second-'0')
	return num
}

func getMatches(text string) [][]int {
	re := regexp.MustCompile(`one|two|thr|fou|fiv|six|sev|eig|nin`)
	m := map[string]int{
		"one": 3,
		"two": 3,
		"thr": 5,
		"fou": 4,
		"fiv": 4,
		"six": 3,
		"sev": 5,
		"eig": 5,
		"nin": 4,
	}
	m2 := map[string]string{
		"one": "one",
		"two": "two",
		"thr": "three",
		"fou": "four",
		"fiv": "five",
		"six": "six",
		"sev": "seven",
		"eig": "eight",
		"nin": "nine",
	}
	index := 0
	matches := make([][]int, 0)
	l := len(text)
	for index < l {
		lastIndex := index + 3
		if lastIndex > len(text) {
			lastIndex = len(text)
		}
		match := re.FindStringSubmatch(text[index:lastIndex])
		if match == nil {
			index++
			continue
		}

		if m2[match[0]] != text[index:index+m[match[0]]] {
			index++
			continue
		}
		matches = append(matches, []int{index, index + m[match[0]]})
		index++
	}
	return matches
}

func readInput(filename string) []string {

	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	open, err := os.Open(filepath.Join(currentDir, filename))
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
