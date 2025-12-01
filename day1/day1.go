package day1

import (
	"strconv"

	aoc "github.com/herman-barnardt/aoc"
)

func init() {
	aoc.Register(2025, 1, solve2025Day1Part1, solve2025Day1Part2)
}

func solve2025Day1Part1(lines []string, test bool) interface{} {
	position := 50
	count := 0
	for _, l := range lines {
		num, _ := strconv.Atoi(l[1:])
		if l[:1] == "L" {
			num = num * -1
		}
		position = (position + num) % 100
		if position < 0 {
			position = 100 + position
		}
		if position == 0 {
			count++
		}
	}
	return count
}

func solve2025Day1Part2(lines []string, test bool) interface{} {
	position := 50
	count := 0
	for _, l := range lines {
		num, _ := strconv.Atoi(l[1:])
		if num > 100 {
			count += num / 100
			num = num % 100
		}
		if l[:1] == "L" {
			num = num * -1
		}
		if position != 0 && ((position+num) >= 100 || (position+num) <= 0) {
			count++
		}
		position = (position + num) % 100
		if position < 0 {
			position = 100 + position
		}
	}
	return count
}
