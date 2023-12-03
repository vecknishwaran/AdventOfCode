package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var part1Sum int
	var part2Sum int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		texts := strings.Split(strings.ReplaceAll(scanner.Text(), " ", ""), ":")

		if IsPossible(texts[1]) {
			part1Sum += getGameNumber(texts[0])
		}

		part2Sum += GetPower(texts[1])
	}
	println("Part 1: ", part1Sum)
	println("Part 2: ", part2Sum)

}

var bag = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func getGameNumber(text string) int {
	s := text[strings.Index(text, "e")+1:]
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func IsPossible(text string) bool {
	set := strings.Split(text, ";")
	for _, colors := range set {
		color := strings.Split(colors, ",")
		for _, c := range color {
			switch {
			case strings.Contains(c, "red"):
				if !isMore(c, "red", bag["red"]) {
					return false
				}
			case strings.Contains(c, "green"):
				if !isMore(c, "green", bag["green"]) {
					return false
				}
			case strings.Contains(c, "blue"):
				if !isMore(c, "blue", bag["blue"]) {
					return false
				}
			}
		}

	}
	return true
}

func GetPower(text string) int {
	set := strings.Split(text, ";")
	maxRed := 0
	maxGreen := 0
	maxBlue := 0
	for _, colors := range set {
		color := strings.Split(colors, ",")
		for _, c := range color {
			switch {
			case strings.Contains(c, "red"):
				maxRed = getMax(c, "red", maxRed)
			case strings.Contains(c, "green"):
				maxGreen = getMax(c, "green", maxGreen)
			case strings.Contains(c, "blue"):
				maxBlue = getMax(c, "blue", maxBlue)
			}
		}

	}

	return maxRed * maxGreen * maxBlue
}

func isMore(text, color string, limit int) bool {
	num := strings.Split(text, color)[0]
	i, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}
	if i > limit {
		return false
	}
	return true
}

func getMax(text, color string, max int) int {
	num := strings.Split(text, color)[0]
	i, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}
	if i > max {
		return i
	}
	return max
}
