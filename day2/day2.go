package day2

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/herman-barnardt/aoc"
)

func init() {
	aoc.Register(2025, 2, solve2025Day2Part1, solve2025Day2Part2)
}

func solve2025Day2Part1(lines []string, test bool) interface{} {
	ranges := strings.Split(lines[0], ",")
	sum := 0
	for _, r := range ranges {
		var start int
		var end int
		fmt.Sscanf(r, "%d-%d", &start, &end)
		for i := start; i <= end; i++ {
			str := strconv.Itoa(i)
			if len(str)%2 == 0 && str[:len(str)/2] == str[len(str)/2:] {
				sum += i
			}
		}

	}
	return sum
}

func solve2025Day2Part2(lines []string, test bool) interface{} {
	ranges := strings.Split(lines[0], ",")
	sum := 0
	for _, r := range ranges {
		var start int
		var end int
		fmt.Sscanf(r, "%d-%d", &start, &end)
		for i := start; i <= end; i++ {
			str := strconv.Itoa(i)
		sequenceLoop:
			for j := 1; j <= len(str); j++ {
				base := str[:j]
				invalid := true
				checked := false
				if len(str)%j == 0 {
					for x := j; x <= len(str)-j; x = x + j {
						invalid = invalid && base == str[x:x+j]
						checked = true
					}
					if invalid && checked {
						sum += i
						break sequenceLoop
					}
				}
			}
		}

	}
	return sum
}
